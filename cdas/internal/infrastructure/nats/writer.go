package nats

import (
	"encoding/xml"
	"fmt"
	"strings"

	"github.com/nats-io/nats.go"
	sfera "github.com/tylland/transition/internal/schema/sfera/v3"
)

// Writer publishes SFERA segment profiles as XML to NATS.
type Writer struct {
	conn *nats.Conn
}

// NewWriter creates a NatsWriter that publishes segment profiles to NATS.
func NewWriter(conn *nats.Conn) *Writer {
	return &Writer{conn: conn}
}

// subjectFromProfile builds the NATS subject from the segment profile.
// Format: sfera.segmentprofile.{imid}.{from_trafikplats}.{to_trafikplats}.{from_nummer}.{to_nummer}.{seq}.{dir}
func subjectFromProfile(segment *sfera.SegmentProfile) (string, error) {
	parts := strings.Split(segment.SPID, ":")
	if len(parts) != 6 {
		return "", fmt.Errorf("unexpected SPID format %q: expected 6 colon-separated parts", segment.SPID)
	}

	var imID *sfera.TeltsiString44

	if segment.SPZone.IMID != nil {
		imID = segment.SPZone.IMID
	}

	// Reorder: trafikplats IDs before node numbers for useful wildcard subscriptions.
	return fmt.Sprintf("sfera.segmentprofile.%d.%s.%s.%s.%s.%s.%s",
		imID, parts[0], parts[2], parts[1], parts[3], parts[4], parts[5]), nil
}

// Write marshals the segment profile to XML and publishes it to its derived NATS subject.
// The message headers carry SP_VersionMajor and SP_VersionMinor from the profile.
// It returns the subject on success.
func (w *Writer) Write(segment *sfera.SegmentProfile) (string, error) {
	subject, err := subjectFromProfile(segment)
	if err != nil {
		return "", err
	}

	data, err := xml.Marshal(segment)
	if err != nil {
		return "", fmt.Errorf("failed to marshal segment profile %s: %w", segment.SPID, err)
	}

	msg := &nats.Msg{
		Subject: subject,
		Data:    data,
		Header: nats.Header{
			"VersionMajor": []string{segment.SPVersionMajor},
			"VersionMinor": []string{*segment.SPVersionMinor},
		},
	}

	if err := w.conn.PublishMsg(msg); err != nil {
		return "", fmt.Errorf("failed to publish segment profile %s to %s: %w", segment.SPID, subject, err)
	}

	return subject, nil
}
