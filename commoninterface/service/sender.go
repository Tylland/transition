package service

import (
	"bytes"
	"crypto/tls"
	"crypto/x509"
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/tylland/transition/commoninterface/message"
	"github.com/tylland/transition/commoninterface/repository"
)

// PartnerConfig holds the bilaterally agreed connection parameters for one CI partner.
// ERA TD 104 §5: partner details must be agreed bilaterally; there is no central registry.
type PartnerConfig struct {
	// CompanyCode is the ERA/UIC organisation code of the partner (e.g. "0000074").
	CompanyCode string
	// EndpointURL is the full HTTPS URL of the partner's inbound web service.
	// ERA TD 104 §5.4.1: https://<CI-IP>:<port>/LIMessageProcessing/http/
	//   UICCCMessageProcessing/UICCCMessageProcessingInboundWS
	EndpointURL string
}

// SenderConfig holds configuration for this CI's outbound transport.
type SenderConfig struct {
	// LocalIP is the public IP address of this CI instance (put into uicmh:messageLiHost).
	LocalIP string
	// Partners maps company code → partner connection details.
	Partners map[string]PartnerConfig
	// TLSCertFile and TLSKeyFile are the PEM paths for mutual TLS (ERA TD 104 §5.4.1).
	// Leave both empty to skip mutual TLS (useful for development/testing).
	TLSCertFile string
	TLSKeyFile  string
	// CACertFile is the PEM path for the CA that signed partner certificates.
	// Leave empty to use the system pool.
	CACertFile string
	// Timeout for each HTTP request.  Defaults to 30 s when zero.
	Timeout time.Duration
	// Validator validates outbound message payloads before sending. Optional; nil skips validation.
	Validator *repository.SchemaRepository
}

// Sender posts TAF/TAP TSI XML messages to partner CI instances via SOAP/HTTPS.
type Sender struct {
	cfg    SenderConfig
	client *http.Client
}

// NewSender creates a Sender.  Call Close (if needed) to release resources.
func NewSender(cfg SenderConfig) (*Sender, error) {
	tlsCfg := &tls.Config{
		MinVersion: tls.VersionTLS12,
	}

	if cfg.TLSCertFile != "" && cfg.TLSKeyFile != "" {
		cert, err := tls.LoadX509KeyPair(cfg.TLSCertFile, cfg.TLSKeyFile)
		if err != nil {
			return nil, fmt.Errorf("ci sender: load client certificate: %w", err)
		}
		tlsCfg.Certificates = []tls.Certificate{cert}
	}

	if cfg.CACertFile != "" {
		pem, err := os.ReadFile(cfg.CACertFile)
		if err != nil {
			return nil, fmt.Errorf("ci sender: read CA cert: %w", err)
		}
		pool := x509.NewCertPool()
		if !pool.AppendCertsFromPEM(pem) {
			return nil, fmt.Errorf("ci sender: no valid certificates in %s", cfg.CACertFile)
		}
		tlsCfg.RootCAs = pool
	}

	timeout := cfg.Timeout
	if timeout == 0 {
		timeout = 30 * time.Second
	}

	return &Sender{
		cfg: cfg,
		client: &http.Client{
			Timeout: timeout,
			Transport: &http.Transport{
				TLSClientConfig: tlsCfg,
			},
		},
	}, nil
}

// SendResult carries the parsed acknowledgement from a partner CI.
type SendResult struct {
	// Status is "ACK" or "NACK".
	Status string
	// AckIdentifier is the unique identifier for this acknowledgement.
	AckIdentifier string
	// ErrorDetail is populated for NACK responses (may be empty).
	ErrorDetail string
}

// Send wraps payload in a SOAP envelope and POSTs it to the partner identified by
// receiverCode.  It returns the parsed acknowledgement or an error.
//
// payload should be a valid TAF/TAP TSI XML document (plain UTF-8).
// msgID should be the TAF/TAP TSI message MessageIdentifier (36-char UUID).
func (s *Sender) Send(receiverCode, msgID, payload string) (SendResult, error) {
	partner, ok := s.cfg.Partners[receiverCode]
	if !ok {
		return SendResult{}, fmt.Errorf("ci sender: unknown partner %q", receiverCode)
	}

	if s.cfg.Validator != nil {
		if err := s.cfg.Validator.Validate(payload); err != nil {
			return SendResult{}, fmt.Errorf("ci sender: %w", err)
		}
	}

	env := message.NewEnvelope(msgID, s.cfg.LocalIP, payload)

	xmlBytes, err := xml.Marshal(env)
	if err != nil {
		return SendResult{}, fmt.Errorf("ci sender: marshal SOAP envelope: %w", err)
	}
	// Prepend the XML declaration.
	body := append([]byte(xml.Header), xmlBytes...)

	req, err := http.NewRequest(http.MethodPost, partner.EndpointURL, bytes.NewReader(body))
	if err != nil {
		return SendResult{}, fmt.Errorf("ci sender: build request: %w", err)
	}
	req.Header.Set("Content-Type", "text/xml; charset=utf-8")
	req.Header.Set("SOAPAction", "")

	resp, err := s.client.Do(req)
	if err != nil {
		return SendResult{}, fmt.Errorf("ci sender: POST to %s: %w", partner.EndpointURL, err)
	}
	defer resp.Body.Close()

	respBytes, err := io.ReadAll(io.LimitReader(resp.Body, 1<<20)) // 1 MiB cap
	if err != nil {
		return SendResult{}, fmt.Errorf("ci sender: read response: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return SendResult{}, fmt.Errorf("ci sender: HTTP %d from %s: %s",
			resp.StatusCode, partner.EndpointURL, respBytes)
	}

	return parseAck(respBytes)
}

// parseAck decodes the SOAP acknowledgement envelope returned by the partner CI.
func parseAck(data []byte) (SendResult, error) {
	var env message.AckEnvelope
	if err := xml.Unmarshal(data, &env); err != nil {
		return SendResult{}, fmt.Errorf("ci sender: unmarshal ACK: %w", err)
	}
	ack := env.Body.Response
	return SendResult{
		Status:        ack.ResponseStatus,
		AckIdentifier: ack.AckIdentifier,
	}, nil
}
