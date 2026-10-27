package jetstream

import (
	"context"
	"fmt"

	"github.com/nats-io/nats.go/jetstream"
)

type Logger struct {
	kv      jetstream.KeyValue
	watcher jetstream.KeyWatcher
}

func NewLogger(js jetstream.JetStream) (*Logger, error) {
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

	watcher, err := kv.WatchAll(ctx)
	if err != nil {
		return nil, err
	}

	return &Logger{kv: kv, watcher: watcher}, nil
}

func (c *Logger) Start(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return
		default:
			for upd := range c.watcher.Updates() {
				if upd != nil {
					fmt.Println(upd.Key())
				} else {
					break
				}
			}
		}
	}
}
