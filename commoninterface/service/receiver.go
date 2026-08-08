package service

import (
	"crypto/tls"
	"crypto/x509"
	"encoding/xml"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/tylland/transition/commoninterface/message"
	"github.com/tylland/transition/commoninterface/repository"
)

// MessageHandler is called by the Receiver for each successfully parsed inbound message.
// msgID is the uicmh:messageIdentifier from the SOAP header.
// payload is the raw TAF/TAP TSI XML from the SOAP body (plain; not yet decompressed/decrypted).
// The returned error causes the receiver to send a NACK; nil sends an ACK.
type MessageHandler func(msgID, senderIP, payload string) error

// ReceiverConfig holds configuration for this CI's inbound transport.
type ReceiverConfig struct {
	// ListenAddr is the TCP address to listen on, e.g. ":443" or "0.0.0.0:8443".
	ListenAddr string
	// TLSCertFile and TLSKeyFile are PEM paths for the server TLS certificate (required).
	TLSCertFile string
	TLSKeyFile  string
	// CACertFile is the PEM path for the CA that signed client certificates.
	// When non-empty, mutual TLS is enforced (ERA TD 104 §5.4.1).
	CACertFile string
	// LIName is the name of this CI instance (max 50 chars), echoed in ACK responses.
	LIName string
	// LIInstanceNumber is the instance number of this CI (max 2 digits), echoed in ACK.
	LIInstanceNumber string
	// Handler processes validated inbound messages.
	Handler MessageHandler
	// Validator validates inbound message payloads. Optional; nil skips validation.
	Validator *repository.SchemaRepository
	// AllowedSenders is an optional set of partner company codes.
	// When non-empty only messages whose uicmh:messageLiHost matches a known partner
	// IP prefix are accepted.  Use nil to accept all TLS-authenticated senders.
	AllowedSenders map[string]PartnerConfig
}

// Receiver listens for inbound SOAP/HTTPS messages from partner CI instances.
// ERA TD 104 §5: the inbound web service path is
//
//	/LIMessageProcessing/http/UICCCMessageProcessing/UICCCMessageProcessingInboundWS
type Receiver struct {
	cfg    ReceiverConfig
	server *http.Server
}

const inboundPath = "/LIMessageProcessing/http/UICCCMessageProcessing/UICCCMessageProcessingInboundWS"
const inboundMessagePattern = "POST " + inboundPath

// NewReceiver creates a Receiver.  Call ListenAndServe to start accepting connections.
func NewReceiver(cfg ReceiverConfig) (*Receiver, error) {
	if cfg.TLSCertFile == "" || cfg.TLSKeyFile == "" {
		return nil, fmt.Errorf("ci receiver: TLSCertFile and TLSKeyFile are required")
	}
	if cfg.Handler == nil {
		return nil, fmt.Errorf("ci receiver: Handler is required")
	}

	tlsCfg := &tls.Config{
		MinVersion: tls.VersionTLS12,
	}

	if cfg.CACertFile != "" {
		pem, err := os.ReadFile(cfg.CACertFile)
		if err != nil {
			return nil, fmt.Errorf("ci receiver: read CA cert: %w", err)
		}
		pool := x509.NewCertPool()
		if !pool.AppendCertsFromPEM(pem) {
			return nil, fmt.Errorf("ci receiver: no valid certificates in %s", cfg.CACertFile)
		}
		tlsCfg.ClientCAs = pool
		tlsCfg.ClientAuth = tls.RequireAndVerifyClientCert
	}

	r := &Receiver{cfg: cfg}
	mux := http.NewServeMux()
	mux.HandleFunc(inboundMessagePattern, r.handleInboundMessage)

	r.server = &http.Server{
		Addr:         cfg.ListenAddr,
		Handler:      mux,
		TLSConfig:    tlsCfg,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
	}
	return r, nil
}

// ListenAndServe starts the HTTPS server.  It blocks until the server shuts down.
func (r *Receiver) ListenAndServe() error {
	log.Printf("ci receiver: listening on %s", r.cfg.ListenAddr)
	return r.server.ListenAndServeTLS(r.cfg.TLSCertFile, r.cfg.TLSKeyFile)
}

// Close shuts down the HTTPS server.
func (r *Receiver) Close() error {
	return r.server.Close()
}

// handleSOAP processes one inbound SOAP request.
func (r *Receiver) handleInboundMessage(w http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	body, err := io.ReadAll(io.LimitReader(req.Body, 10<<20)) // 10 MiB cap
	if err != nil {
		log.Printf("ci receiver: read body: %v", err)
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	env, err := parseEnvelope(body)
	if err != nil {
		log.Printf("ci receiver: parse SOAP envelope: %v", err)
		writeAck(w, r.cfg, "", "", "", "NACK")
		return
	}

	msgID := env.Header.MessageIdentifier
	senderIP := env.Header.MessageLiHost
	payload := env.Body.UICMessage.Message

	if r.cfg.Validator != nil {
		if err := r.cfg.Validator.Validate(payload); err != nil {
			log.Printf("ci receiver: schema validation failed for message %s: %v", msgID, err)
			writeAck(w, r.cfg, msgID, senderIP, payload, "NACK")
			return
		}
	}

	// Call the application-level handler.
	if handlerErr := r.cfg.Handler(msgID, senderIP, payload); handlerErr != nil {
		log.Printf("ci receiver: handler rejected message %s: %v", msgID, handlerErr)
		writeAck(w, r.cfg, msgID, senderIP, payload, "NACK")
		return
	}

	writeAck(w, r.cfg, msgID, senderIP, payload, "ACK")
}

// parseEnvelope decodes the SOAP request body.
func parseEnvelope(data []byte) (message.Envelope, error) {
	var env message.Envelope
	if err := xml.Unmarshal(data, &env); err != nil {
		return message.Envelope{}, err
	}
	return env, nil
}

// writeAck sends a SOAP ACK or NACK response back to the caller.
func writeAck(w http.ResponseWriter, cfg ReceiverConfig, msgID, senderIP, _ string, status string) {
	ackID := "ACKID" + msgID

	ack := message.AckResponse{
		ResponseStatus:            status,
		AckIdentifier:             ackID,
		MessageReference:          message.MessageReference{MessageIdentifier: msgID},
		RemoteLIName:              cfg.LIName,
		RemoteLIInstanceNumber:    cfg.LIInstanceNumber,
		MessageTransportMechanism: "WEBSERVICE",
	}

	respEnv := buildAckEnvelope(ack)
	xmlBytes, err := xml.MarshalIndent(respEnv, "", "  ")
	if err != nil {
		log.Printf("ci receiver: marshal ACK: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/xml; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte(xml.Header))
	_, _ = w.Write(xmlBytes)
}

// buildAckEnvelope wraps an AckResponse in a SOAP envelope.
func buildAckEnvelope(ack message.AckResponse) outAckEnvelope {
	return outAckEnvelope{
		NsSOAP: message.NsSOAP,
		Body:   outAckBody{Response: ack},
	}
}

// outAckEnvelope is used only for marshalling ACK responses.
type outAckEnvelope struct {
	XMLName xml.Name   `xml:"soapenv:Envelope"`
	NsSOAP  string     `xml:"xmlns:soapenv,attr"`
	Body    outAckBody `xml:"soapenv:Body"`
}

type outAckBody struct {
	Response message.AckResponse `xml:"response"`
}

// WSDLHandler returns an http.HandlerFunc that serves a minimal WSDL description
// for the inbound web service (ERA TD 104 §5.1 / §5.4.1).
func WSDLHandler(baseURL string) http.HandlerFunc {
	wsdl := buildWSDL(baseURL)
	return func(w http.ResponseWriter, req *http.Request) {
		if !strings.HasSuffix(req.URL.RawQuery, "wsdl") &&
			!strings.EqualFold(req.URL.Query().Get("wsdl"), "") {
			http.NotFound(w, req)
			return
		}
		w.Header().Set("Content-Type", "text/xml; charset=utf-8")
		_, _ = w.Write([]byte(wsdl))
	}
}

// buildWSDL returns a minimal WSDL 1.1 document for the inbound service.
func buildWSDL(baseURL string) string {
	return `<?xml version="1.0" encoding="UTF-8"?>
<definitions xmlns="http://schemas.xmlsoap.org/wsdl/"
             xmlns:soap="http://schemas.xmlsoap.org/wsdl/soap/"
             xmlns:tns="http://uic.cc.org/UICMessage"
             xmlns:uicm="http://uic.cc.org/UICMessage"
             targetNamespace="http://uic.cc.org/UICMessage"
             name="UICCCMessageProcessingInboundWS">
  <types>
    <schema xmlns="http://www.w3.org/2001/XMLSchema"
            targetNamespace="http://uic.cc.org/UICMessage">
      <element name="UICMessage">
        <complexType>
          <sequence>
            <element name="message"      type="string"/>
            <element name="signature"    type="string" minOccurs="0"/>
            <element name="senderAlias"  type="string" minOccurs="0"/>
            <element name="encoding"     type="string"/>
          </sequence>
        </complexType>
      </element>
    </schema>
  </types>
  <message name="sendUICMessageRequest">
    <part name="parameters" element="uicm:UICMessage"/>
  </message>
  <message name="sendUICMessageResponse">
    <part name="parameters" element="uicm:UICMessage"/>
  </message>
  <portType name="UICCCMessageProcessingInboundWSPortType">
    <operation name="sendUICMessage">
      <input  message="tns:sendUICMessageRequest"/>
      <output message="tns:sendUICMessageResponse"/>
    </operation>
  </portType>
  <binding name="UICCCMessageProcessingInboundWSBinding"
           type="tns:UICCCMessageProcessingInboundWSPortType">
    <soap:binding style="document"
                  transport="http://schemas.xmlsoap.org/soap/http"/>
    <operation name="sendUICMessage">
      <soap:operation soapAction=""/>
      <input><soap:body use="literal"/></input>
      <output><soap:body use="literal"/></output>
    </operation>
  </binding>
  <service name="UICCCMessageProcessingInboundWS">
    <port name="UICCCMessageProcessingInboundWSPort"
          binding="tns:UICCCMessageProcessingInboundWSBinding">
      <soap:address location="` + baseURL + inboundPath + `"/>
    </port>
  </service>
</definitions>`
}
