package message

import (
	"encoding/xml"
	"fmt"
	"os"
	"testing"

	"github.com/tylland/transition/specs"
)

const (
	validPayload     = `<TrainRunningInformationMessage><MessageHeader/></TrainRunningInformationMessage>`
	wrongRootPayload = `<PathRequestMessage><Data/></PathRequestMessage>`
	malformedPayload = `<NotClosed`
)

func TestEnvelopeRoundtrip(t *testing.T) {
	senv := NewEnvelope("1123", "99", validPayload)

	data, _ := xml.MarshalIndent(senv, "", "  ")

	fmt.Print(string(data))

	var renv Envelope

	err := xml.Unmarshal(data, &renv)
	if err != nil {
		fmt.Print(string(data))
		t.Fatal(err)
	}

	if renv.Header.MessageIdentifier != senv.Header.MessageIdentifier {
		fmt.Print(string(data))
		t.Fatalf("MessageIdentifier expected %q got %q", senv.Header.MessageIdentifier, renv.Header.MessageIdentifier)
	}

	if renv.Header.MessageLiHost != senv.Header.MessageLiHost {
		fmt.Print(string(data))
		t.Fatalf("MessageLiHost expected %q got %q", senv.Header.MessageLiHost, renv.Header.MessageLiHost)
	}

	// if renv.Body.UICMessage.Message != senv.Body.UICMessage.Message {
	// 	fmt.Print(string(data))
	// 	t.Fatalf("UICMessage expected %q got %q", senv.Body.UICMessage.Message, renv.Body.UICMessage.Message)
	// }
}

func TestAckEnvelopeRoundtrip(t *testing.T) {
	sack := AckResponse{ResponseStatus: ResponseStatusAccepted, AckIdentifier: "ACKID-test"}
	sackenv := NewAckEnvelope(sack)

	data, _ := xml.MarshalIndent(sackenv, "", "")

	var rackenv AckEnvelope
	xml.Unmarshal(data, &rackenv)

	rack := rackenv.Body.Response

	if rack.ResponseStatus != sack.ResponseStatus {
		t.Fatalf("MessageIdentifier expected %q got %q", sack.ResponseStatus, rack.ResponseStatus)
	}

	if rack.AckIdentifier != sack.AckIdentifier {
		t.Fatalf("MessageIdentifier expected %q got %q", sack.AckIdentifier, rack.AckIdentifier)
	}

}

func TestReadTrainComposition(t *testing.T) {

	data := specs.Must(os.ReadFile("../testdata/traincomposition.xml"))

	var renv Envelope

	err := xml.Unmarshal(data, &renv)
	if err != nil {
		fmt.Print(string(data))
		t.Fatal(err)
	}

	messageIdentifier := "796cc4ee-cd51-49c2-a8cf-3bd73b29ddfa"

	if renv.Header.MessageIdentifier != messageIdentifier {
		fmt.Print(string(data))
		t.Fatalf("MessageIdentifier expected %q got %q", messageIdentifier, renv.Header.MessageIdentifier)
	}

	// if renv.Header.MessageIdentifier != senv.Header.MessageIdentifier {
	// 	fmt.Print(string(data))
	// 	t.Fatalf("MessageIdentifier expected %q got %q", senv.Header.MessageIdentifier, renv.Header.MessageIdentifier)
	// }

	// if renv.Header.MessageLiHost != senv.Header.MessageLiHost {
	// 	fmt.Print(string(data))
	// 	t.Fatalf("MessageLiHost expected %q got %q", senv.Header.MessageLiHost, renv.Header.MessageLiHost)
	// }

	// if renv.Body.UICMessage.Message != senv.Body.UICMessage.Message {
	// 	fmt.Print(string(data))
	// 	t.Fatalf("UICMessage expected %q got %q", senv.Body.UICMessage.Message, renv.Body.UICMessage.Message)
	// }
}
