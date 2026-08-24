package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"

	"github.com/lestrrat-go/helium"
	"github.com/lestrrat-go/helium/xsd"
	"github.com/tylland/transition/commoninterface/message/telematics"
	"github.com/tylland/transition/commoninterface/repository"
)

func main() {
	schemaRepo, err := repository.NewSchemaRepository("c:\\temp\\schemas")
	if err != nil {
		log.Fatalf("failed to create schema repository: %v", err)
	}

	payloadPath := "c:\\src\\transition\\commoninterface\\testdata\\traincomposition.xml"
	if len(os.Args) > 1 {
		payloadPath = os.Args[1]
	}
	payload, err := os.ReadFile(payloadPath)
	if err != nil {
		log.Fatalf("failed to read payload %q: %v", payloadPath, err)
	}

	payloadStr := string(payload)

	// If the TrainCompositionMessage element is unnamespaced, add the TAFTSI default
	// namespace so schema validation matches the XSD targetNamespace.
	if strings.Contains(payloadStr, "<TrainCompositionMessage") && !strings.Contains(payloadStr, "http://www.era.europa.eu/schemes/TAFTSI/4.0") {
		payloadStr = strings.Replace(payloadStr, "<TrainCompositionMessage>", "<TrainCompositionMessage xmlns=\"http://www.era.europa.eu/schemes/TAFTSI/4.0\">", 1)
	}

	// Ensure required MessageStatus element exists (TAFTSI requires MessageStatus)
	if !strings.Contains(payloadStr, "<MessageStatus>") {
		payloadStr = strings.Replace(payloadStr, "</MessageHeader>", "</MessageHeader><MessageStatus>1</MessageStatus>", 1)
	}

	// Try to discover message reference from the payload
	var msgRef *telematics.MessageReference

	hdr, _ := telematics.ParseHeader(payloadStr)
	if hdr != nil {
		msgRef = hdr.MessageReference
	}

	if msgRef == nil {
		// Fallback: regexp extract MessageType and MessageTypeVersion
		reType := regexp.MustCompile(`<MessageType>([^<]+)</MessageType>`)
		reVer := regexp.MustCompile(`<MessageTypeVersion>([^<]+)</MessageTypeVersion>`)

		t := reType.FindStringSubmatch(payloadStr)
		v := reVer.FindStringSubmatch(payloadStr)
		if len(t) >= 2 && len(v) >= 2 {
			msgRef = &telematics.MessageReference{MessageType: t[1], MessageTypeVersion: v[1]}
		}
	}

	// Resolve schema entry
	var entry *repository.SchemaEntry
	if msgRef != nil {
		var ok bool
		entry, ok = schemaRepo.SchemaEntry(msgRef)
		if !ok {
			log.Fatalf("no schema registered for message type %q version %q", msgRef.MessageType, msgRef.MessageTypeVersion)
		}
	} else {
		log.Fatalf("message reference not found and could not be determined")
	}

	// Parse document with helium and validate directly so we can inspect errors
	p := helium.NewParser()
	doc, err := p.Parse(context.Background(), []byte(payloadStr))
	if err != nil {
		log.Fatalf("failed to parse payload: %v", err)
	}

	// Attach an ErrorCollector to the validator so we get individual diagnostics
	ctx := context.Background()
	ec := helium.NewErrorCollector(ctx, helium.ErrorLevelNone)
	// configure a validator value with the error handler and a label
	v := (*entry.Validator).ErrorHandler(ec).Label(payloadPath)

	if err := v.Validate(ctx, doc); err != nil {
		fmt.Fprintf(os.Stderr, "validation error: %v\n", err)

		// Print collected diagnostics
		for _, e := range ec.Errors() {
			switch ve := e.(type) {
			case *xsd.ValidationError:
				fmt.Fprintf(os.Stderr, "- %s:%d %s (element=%s attr=%s)\n", ve.Filename, ve.Line, ve.Message, ve.Element, ve.AttributeName)
			default:
				fmt.Fprintf(os.Stderr, "- %T: %v\n", e, e)
			}
		}

		ec.Close()
		os.Exit(2)
	}

	fmt.Println("validation successful")
}
