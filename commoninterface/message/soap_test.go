package message

import (
	"encoding/xml"
	"testing"
)

const (
	validPayload     = `<TrainRunningInformationMessage><MessageHeader/></TrainRunningInformationMessage>`
	wrongRootPayload = `<PathRequestMessage><Data/></PathRequestMessage>`
	malformedPayload = `<NotClosed`
)

func TestRoundtrip(t *testing.T) {
	senv := NewEnvelope("1123", "99", validPayload)

	data, _ := xml.MarshalIndent(senv, "", "")

	var renv Envelope

	xml.Unmarshal(data, &renv)
}
