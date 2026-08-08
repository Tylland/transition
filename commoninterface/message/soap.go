// Package ci implements the ERA TD 104 Common Interface transport layer.
// CI-to-CI communication uses SOAP over HTTPS (push model).
// Reference: ERA_TD_104 v4.0, Chapter 5.
package message

import "encoding/xml"

// Namespaces defined in ERA TD 104 §5.2.
const (
	NsSOAP   = "http://schemas.xmlsoap.org/soap/envelope/"
	NsHeader = "http://uic.cc.org/UICMessage/Header"
	NsBody   = "http://uic.cc.org/UICMessage"
)

// Envelope is a SOAP 1.1 envelope (ERA TD 104 §5.2).
type Envelope struct {
	XMLName xml.Name `xml:"http://schemas.xmlsoap.org/soap/envelope/ Envelope"`
	//XMLName xml.Name `xml:"soapenv:Envelope"`
	NsSOAP  string `xml:"xmlns:soapenv,attr"`
	NsUICMH string `xml:"xmlns:uicmh,attr"`
	NsUICM  string `xml:"xmlns:uicm,attr"`
	Header  Header
	Body    Body
}

// Header carries the CI routing metadata (ERA TD 104 §5.2.1).
type Header struct {
	// MessageIdentifier uniquely identifies the message (16-byte GUID, 36-char string).
	MessageIdentifier string `xml:"messageIdentifier"`
	// Compressed indicates whether the payload in Body.UICMessage.Message is zlib-compressed.
	Compressed bool `xml:"uicmh:compressed"`
	// Encrypted indicates whether the payload is encrypted.
	Encrypted bool `xml:"uicmh:encrypted"`
	// Signed indicates whether the payload carries a digital signature.
	Signed bool `xml:"uicmh:signed"`
	// MessageLiHost is the public IP address of the sending CI instance.
	MessageLiHost string `xml:"messageLiHost"`
}

// Body wraps the UICMessage payload.
type Body struct {
	UICMessage UICMessage `xml:"UICMessage"`
}

// UICMessage is the SOAP body payload (ERA TD 104 §5.2.2).
type UICMessage struct {
	// Message is the TAF/TAP TSI XML payload.
	// If compressed, encrypted, or signed it is base64-encoded.
	Message string `xml:"uicm:message"`
	// Signature is the optional digital signature (SHA256withECDSA).
	Signature string `xml:"uicm:signature,omitempty"`
	// SenderAlias is the alias of the sender's TLS certificate in the receiver's truststore.
	SenderAlias string `xml:"uicm:senderAlias,omitempty"`
	// Encoding is the character encoding of the message (e.g. "UTF-8").
	Encoding string `xml:"uicm:encoding"`
}

// AckEnvelope is the SOAP response envelope containing the acknowledgement.
type AckEnvelope struct {
	XMLName xml.Name `xml:"Envelope"`
	Body    AckBody  `xml:"Body"`
}

// AckBody wraps the acknowledgement XML.
type AckBody struct {
	Response AckResponse `xml:"response"`
}

// AckResponse is the acknowledgement body (ERA TD 104 §5.3).
type AckResponse struct {
	// ResponseStatus is "ACK" (accepted) or "NACK" (rejected).
	ResponseStatus string `xml:"ResponseStatus"`
	// AckIdentifier is prefixed with "ACKID" followed by the original MessageIdentifier.
	AckIdentifier string `xml:"AckIndentifier"` // sic: spec uses "AckIndentifier"
	// MessageReference echoes key fields from the received message header.
	MessageReference MessageReference `xml:"MessageReference"`
	Sender           string           `xml:"Sender"`
	// Recipient is the receiver company ID found in the TAF/TAP TSI message header.
	Recipient string `xml:"Recipient"`
	// RemoteLIName is the name of the receiving CI (max 50 chars).
	RemoteLIName string `xml:"RemoteLIName"`
	// RemoteLIInstanceNumber is the instance number of the receiving CI (max 2 digits).
	RemoteLIInstanceNumber string `xml:"RemoteLIInstanceNumber"`
	// MessageTransportMechanism is always "WEBSERVICE" for SOAP transport.
	MessageTransportMechanism string `xml:"MessageTransportMechanism"`
}

// MessageReference echoes metadata from the received message (ERA TD 104 §5.3).
type MessageReference struct {
	MessageType        string `xml:"MessageType"`
	MessageTypeVersion string `xml:"MessageTypeVersion"`
	MessageIdentifier  string `xml:"MessageIdentifier"`
	MessageDateTime    string `xml:"MessageDateTime"`
}

// NewEnvelope constructs a SOAP envelope ready to send.
// payload is the raw TAF/TAP TSI XML (plain, uncompressed, unencrypted).
func NewEnvelope(msgID, senderIP, payload string) Envelope {
	return Envelope{
		NsSOAP:  NsSOAP,
		NsUICMH: NsHeader,
		NsUICM:  NsBody,
		Header: Header{
			MessageIdentifier: msgID,
			Compressed:        false,
			Encrypted:         false,
			Signed:            false,
			MessageLiHost:     senderIP,
		},
		Body: Body{
			UICMessage: UICMessage{
				Message:  payload,
				Encoding: "UTF-8",
			},
		},
	}
}
