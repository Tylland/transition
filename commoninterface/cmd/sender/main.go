// Command sender is a CLI tool for sending a TAF/TAP TSI XML message
// to a partner Common Interface via SOAP/HTTPS (ERA TD 104 §5).
//
// Usage:
//
//	sender -to <company-code> -msg <file.xml> [flags]
//
// Flags:
//
//	-to        Partner ERA/UIC company code (required)
//	-url       Partner endpoint URL (required)
//	-id        Message identifier (UUID); generated if omitted
//	-msg       Path to TAF/TAP TSI XML file to send (required)
//	-local-ip  Public IP of this CI instance (put in uicmh:messageLiHost)
//	-cert      PEM path for mutual TLS client certificate
//	-key       PEM path for mutual TLS client key
//	-ca        PEM path for CA certificate (optional)
package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/google/uuid"
	ci "github.com/tylland/transition/commoninterface/service"
)

func main() {
	to := flag.String("to", "", "Partner company code (e.g. 0000074)")
	partnerURL := flag.String("url", "", "Partner endpoint URL")
	msgID := flag.String("id", "", "Message identifier (UUID); auto-generated if empty")
	msgFile := flag.String("msg", "", "Path to TAF/TAP TSI XML file to send")
	localIP := flag.String("local-ip", "127.0.0.1", "Public IP of this CI instance")
	certFile := flag.String("cert", "", "PEM path for mutual TLS client certificate")
	keyFile := flag.String("key", "", "PEM path for mutual TLS client key")
	caFile := flag.String("ca", "", "PEM path for CA certificate (optional)")
	flag.Parse()

	if *to == "" || *partnerURL == "" || *msgFile == "" {
		flag.Usage()
		os.Exit(1)
	}

	if *msgID == "" {
		*msgID = uuid.New().String()
	}

	payload, err := os.ReadFile(*msgFile)
	if err != nil {
		log.Fatalf("read message file: %v", err)
	}

	cfg := ci.SenderConfig{
		LocalIP: *localIP,
		Partners: map[string]ci.PartnerConfig{
			*to: {CompanyCode: *to, EndpointURL: *partnerURL},
		},
		TLSCertFile: *certFile,
		TLSKeyFile:  *keyFile,
		CACertFile:  *caFile,
	}

	sender, err := ci.NewSender(cfg)
	if err != nil {
		log.Fatalf("create sender: %v", err)
	}

	result, err := sender.Send(*to, *msgID, string(payload))
	if err != nil {
		log.Fatalf("send: %v", err)
	}

	fmt.Printf("status:        %s\n", result.Status)
	fmt.Printf("ack-identifier: %s\n", result.AckIdentifier)
	if result.Status != "ACK" {
		os.Exit(2)
	}
}
