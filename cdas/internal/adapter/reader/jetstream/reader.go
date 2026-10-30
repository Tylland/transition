package jetstream

import (
	"context"
	"encoding/xml"

	"github.com/nats-io/nats.go/jetstream"
	sfera "github.com/tylland/transition/cdas/internal/schema/sfera/v3"
)

type KeyValueReader struct {
	kv      jetstream.KeyValue
	watcher jetstream.KeyWatcher
}

func NewKeyValueReader(js jetstream.JetStream, bucket string) (*KeyValueReader, error) {
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

	return &KeyValueReader{kv: kv}, nil
}

func (r *KeyValueReader) Read[T any](ctx context.Context, key string) (*T, error) {
	ent, err := r.kv.Get(ctx, key)

	if err != nil {
		return nil, err
	}

	v := new(T)

	err = xml.Unmarshal(ent.Value(), v)

	if err != nil {
		return nil, err
	}

	return v, nil
}

type SegmentProfileReader struct {
	*KeyValueReader
}

func (r *SegmentProfileReader) Read(ctx context.Context, key string) (*sfera.SegmentProfile, error) {
	ent, err := r.kv.Get(ctx, key)

	segment := &sfera.SegmentProfile{}

	err = xml.Unmarshal(ent.Value(), segment)

	if err != nil {
		return nil, err
	}

	return segment, nil
}
