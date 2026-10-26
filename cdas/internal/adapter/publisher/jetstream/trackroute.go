package jetstream

import (
	"context"
	"encoding/xml"
	"fmt"

	"github.com/nats-io/nats.go/jetstream"
	"github.com/tylland/transition/cdas/internal/core/domain"
)

type TrackRouteKVWriter struct {
	kv jetstream.KeyValue
}

func NewTrackRouteKVWriter(js jetstream.JetStream) (*TrackRouteKVWriter, error) {
	const bucket = "TRACK_ROUTES"

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

	return &TrackRouteKVWriter{kv: kv}, nil
}

func (w *TrackRouteKVWriter) key(route *domain.TrackRoute) string {
	//	return fmt.Sprintf("%s:%s:%s:%s", route.StartPlace, route.EndPlace, route.StartSignal, route.EndSignal), nil
	return fmt.Sprintf("%s", route.Id)
}

func (w *TrackRouteKVWriter) Write(routes *domain.TrackRoutes) error {
	ctx := context.Background()

	for _, route := range routes.Routes {
		key := w.key(route)

		data, err := xml.Marshal(route)
		if err != nil {
			return fmt.Errorf("failed to marshal signal route %s: %w", key, err)
		}

		if _, err := w.kv.Put(ctx, key, data); err != nil {
			return fmt.Errorf("save %q to bucket: %w", key, err)
		}
	}

	return nil
}
