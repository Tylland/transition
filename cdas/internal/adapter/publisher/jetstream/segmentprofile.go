package jetstream

import (
	"context"
	"encoding/xml"
	"fmt"

	"github.com/nats-io/nats.go/jetstream"
	sfera "github.com/tylland/transition/cdas/internal/schema/sfera/v3"
)

type SegmentProfileKVWriter struct {
	kv jetstream.KeyValue
}

func NewSegmentProfileKVWriter(js jetstream.JetStream) (*SegmentProfileKVWriter, error) {
	const bucket = "SEGMENT_PROFILES"

	ctx := context.Background()

	kv, err := js.KeyValue(ctx, bucket)
	if err != nil {
		kv, err = js.CreateKeyValue(ctx, jetstream.KeyValueConfig{
			Bucket: bucket,
		})
		if err != nil {
			return nil, err
		}
	}

	return &SegmentProfileKVWriter{kv: kv}, nil
}

// versionedKey builds the NATS key from the segment profile.
// Format: sfera.segmentprofile.{imid}.{from_place}.{to_place}.{from_node}.{to_node}.{seq}.{dir}.{major}.{minor}
func versionedKey(segment *sfera.SegmentProfile) (string, error) {
	major := segment.SPVersionMajor
	minor := "0"

	if segment.SPVersionMinor != nil {
		minor = *segment.SPVersionMinor
	}

	return fmt.Sprintf("%s.%s.%s",
		segment.SPID, major, minor), nil
}

// Write marshals the segment profile to XML and publishes it.
// It returns the key on success.
func (w *SegmentProfileKVWriter) Write(segment *sfera.SegmentProfile) (string, error) {
	ctx := context.Background()

	key, err := versionedKey(segment)
	if err != nil {
		return "", err
	}

	data, err := xml.Marshal(segment)
	if err != nil {
		return "", fmt.Errorf("failed to marshal segment profile %s: %w", segment.SPID, err)
	}

	if _, err := w.kv.Put(ctx, key, data); err != nil {
		return "", fmt.Errorf("save %q to bucket: %w", key, err)
	}

	currentKey := "current." + segment.SPID
	_, err = w.kv.Put(ctx, currentKey, []byte(key))

	return key, nil
}
