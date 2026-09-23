// Command receiver runs a CI inbound SOAP/HTTPS web service (ERA TD 104 §5).
// It listens for TAF/TAP TSI XML messages pushed by partner CI instances and
// writes each accepted payload to a file in the configured output directory.
//
// Usage:
//
//	receiver [flags]
//
// Flags:
//
//	-addr       TCP listen address (default :8443)
//	-cert       PEM path for server TLS certificate (required)
//	-key        PEM path for server TLS key (required)
//	-ca         PEM path for client CA certificate (enables mutual TLS)
//	-out        Directory to write received XML messages (default ./received)
//	-li-name    Name of this CI instance (echoed in ACK, max 50 chars)
//	-li-inst    Instance number of this CI (echoed in ACK, max 2 digits)
package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sync"
	"time"

	"github.com/tylland/transition/commoninterface/repository"
	"github.com/tylland/transition/commoninterface/service"
)

func main() {
	addr := flag.String("addr", ":8443", "TCP listen address")
	certFile := flag.String("cert", "", "PEM path for server TLS certificate (required)")
	keyFile := flag.String("key", "", "PEM path for server TLS key (required)")
	caFile := flag.String("ca", "", "PEM path for client CA certificate (mutual TLS)")
	outDir := flag.String("out", "received", "Directory to write received XML messages")
	liName := flag.String("li-name", "LocalCI", "Name of this CI instance (max 50 chars)")
	liInst := flag.String("li-inst", "01", "Instance number of this CI (max 2 digits)")
	flag.Parse()

	if *certFile == "" || *keyFile == "" {
		flag.Usage()
		os.Exit(1)
	}

	if err := os.MkdirAll(*outDir, 0o700); err != nil {
		log.Fatalf("create output directory %s: %v", *outDir, err)
	}

	schemaRepo, err := repository.NewSchemaRepository("C:\\src\\transition\\commoninterface\\data\\reposotory\\schemas")

	if err != nil {
		log.Fatalf("create output directory %s: %v", *outDir, err)
	}

	handler := func(msgID, senderIP, payload string) error {
		filename := filepath.Join(*outDir, fmt.Sprintf("%s_%s.xml",
			time.Now().UTC().Format("20060102T150405Z"),
			msgID))
		if err := os.WriteFile(filename, []byte(payload), 0o600); err != nil {
			return fmt.Errorf("write payload to %s: %w", filename, err)
		}
		log.Printf("received message id=%q from=%q saved=%q", msgID, senderIP, filename)
		return nil
	}

	cfg := service.ReceiverConfig{
		ListenAddr:       *addr,
		TLSCertFile:      *certFile,
		TLSKeyFile:       *keyFile,
		CACertFile:       *caFile,
		LIName:           *liName,
		LIInstanceNumber: *liInst,
		Handler:          handler,
		Validator:        schemaRepo,
	}

	recv, err := service.NewReceiver(cfg)
	if err != nil {
		log.Fatalf("create receiver: %v", err)
	}

	defer recv.Close()

	wg := sync.WaitGroup{}

	wg.Add(2)

	go StartSecureServer(recv, &wg)
	go StartUnsafeServer(recv, &wg)

	wg.Wait()
}

func StartSecureServer(recv *service.Receiver, wg *sync.WaitGroup) {
	defer wg.Done()

	if err := recv.StartSecureServer(); err != nil {
		log.Fatalf("receiver: %v", err)
	}
}

func StartUnsafeServer(recv *service.Receiver, wg *sync.WaitGroup) {
	defer wg.Done()

	if err := recv.StartUnsafeServer(); err != nil {
		log.Fatalf("receiver: %v", err)
	}

}
