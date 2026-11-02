package jetstream

// type ActiveTrain struct {
// 	js jetstream.JetStream
// }

// func NewActiveTrain(js jetstream.JetStream) (*Logger, error) {

// 	ctx := context.Background()

// 	kv, err := js.KeyValue(ctx, bucket)
// 	if err != nil {
// 		kv, err = js.CreateKeyValue(ctx, jetstream.KeyValueConfig{
// 			Bucket: bucket,
// 		})
// 		if err != nil {
// 			return nil, err
// 		}
// 	}

// 	watcher, err := kv.WatchAll(ctx)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return &Logger{kv: kv, watcher: watcher}, nil
// }

// func (c *ActiveTrain) StartSession() {

// }
