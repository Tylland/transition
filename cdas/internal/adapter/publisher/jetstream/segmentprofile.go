package jetstream

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"encoding/xml"
	"fmt"
	"os"
	"strconv"

	"github.com/nats-io/nats.go/jetstream"
	"github.com/tylland/transition/cdas/internal/core/domain"
	"github.com/tylland/transition/cdas/internal/core/port"
	sfera "github.com/tylland/transition/cdas/internal/schema/sfera/v3"
)

type SegmentProfileKVWriter struct {
	key port.SPKeyProvider
	kv  jetstream.KeyValue
}

func NewSegmentProfileKVWriter(js jetstream.JetStream, key port.SPKeyProvider) (*SegmentProfileKVWriter, error) {
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

	return &SegmentProfileKVWriter{key: key, kv: kv}, nil
}

// versionedKey builds the NATS key from the segment profile.
// Format: sfera.segmentprofile.{imid}.{from_place}.{to_place}.{from_node}.{to_node}.{seq}.{dir}.{major}.{minor}
func versionedKey(segment *sfera.SegmentProfile) (string, error) {
	major := segment.SPVersionMajor
	minor := "0"

	if segment.SPVersionMinor != nil {
		minor = *segment.SPVersionMinor
	}

	return fmt.Sprintf("sfera.segmentprofile.%s.%s.%s",
		segment.SPID, major, minor), nil
}

// hashSHA256 takes a string and returns its SHA-256 hash in hex format
func hashSHA256(data []byte) string {
	// Create a new SHA-256 hash object
	hash := sha256.New()

	// Write the input data to the hash
	_, err := hash.Write(data)
	if err != nil {
		// This should rarely happen, but handle it gracefully
		fmt.Fprintf(os.Stderr, "Error writing to hash: %v\n", err)
		return ""
	}

	// Compute the final hash sum
	hashBytes := hash.Sum(nil)

	// Convert the hash bytes to a hexadecimal string
	return hex.EncodeToString(hashBytes)
}

func checksum(seg *sfera.SegmentProfile) (string, error) {
	checksumData := &SegmentProfileChecksumData{
		SPLength:              seg.SPLength,
		SPName:                seg.SPName,
		SPStatus:              seg.SPStatus,
		DistanceEoAOffset:     seg.DistanceEoAOffset,
		UTCOffset:             seg.UTCOffset,
		SPAltitude:            seg.SPAltitude,
		SPGauge:               seg.SPGauge,
		SPMaxLateralDeviation: seg.SPMaxLateralDeviation,
		SPPoints:              seg.SPPoints,
		SPAreas:               seg.SPAreas,
		SPCharacteristics:     seg.SPCharacteristics,
		SPContextInformation:  seg.SPContextInformation,
	}

	data, err := xml.Marshal(checksumData)
	if err != nil {
		return "", fmt.Errorf("failed to marshal segment profile %s: %w", seg.SPID, err)
	}

	return hashSHA256(data), nil
}

// Write marshals the segment profile to XML and publishes it.
// It returns the key on success.
func (w *SegmentProfileKVWriter) Write(segment *sfera.SegmentProfile) (string, error) {
	ctx := context.Background()

	key := w.key.Data(segment.SPID)
	var major int = 1
	var minor int = 1

	checksum, err := checksum(segment)
	if err != nil {
		return "", fmt.Errorf("creating checksum %s: %w", segment.SPID, err)
	}

	infoKey := w.key.Info(segment.SPID)
	infoEntry, err := w.kv.Get(ctx, infoKey)

	if infoEntry != nil {
		info := &domain.SegmentProfileInfo{}
		err = xml.Unmarshal(infoEntry.Value(), info)

		if err != nil {
			return "", fmt.Errorf("umarshal segment profile info %s: %w", segment.SPID, err)
		}

		if info.Checksum == checksum {
			return key, nil
		}

		major = info.Major + 1
	}

	segment.SPVersionMajor = strconv.Itoa(major)
	segment.SPVersionMinor = new(strconv.Itoa(minor))

	data, err := xml.Marshal(segment)
	if err != nil {
		return "", fmt.Errorf("marshal segment profile %s: %w", segment.SPID, err)
	}

	_, err = w.kv.Put(ctx, key, data)

	if err != nil {
		return "", fmt.Errorf("save %q to bucket: %w", key, err)
	}

	info := &domain.SegmentProfileInfo{Major: major, Minor: 0, Checksum: checksum, Segment: segment}

	infoData, err := xml.Marshal(info)
	if err != nil {
		return "", fmt.Errorf("marshal info %q: %w", segment.SPID, err)
	}

	_, err = w.kv.Put(ctx, infoKey, infoData)
	if err != nil {
		return "", fmt.Errorf("save info %q: %w", segment.SPID, err)
	}

	return key, nil
}

type SegmentProfileChecksumData struct {
	SPLength              float64                     `xml:"SP_Length,attr"`
	SPName                *string                     `xml:"SP_Name,attr"`
	SPStatus              sfera.SPStatus              `xml:"SP_Status,attr"`
	DistanceEoAOffset     *float64                    `xml:"distance_EoA_Offset,attr"`
	UTCOffset             *uint8                      `xml:"UTC_Offset,attr"`
	SPAltitude            *float64                    `xml:"SP_Altitude,attr"`
	SPGauge               *int                        `xml:"SP_Gauge,attr"`
	SPMaxLateralDeviation *float64                    `xml:"SP_MaxLateralDeviation,attr"`
	SPPoints              *sfera.SPPoints             `xml:"SP_Points"`
	SPAreas               *sfera.SPAreas              `xml:"SP_Areas"`
	SPCharacteristics     *sfera.SPCharacteristics    `xml:"SP_Characteristics"`
	SPContextInformation  *sfera.SPContextInformation `xml:"SP_ContextInformation"`
}
