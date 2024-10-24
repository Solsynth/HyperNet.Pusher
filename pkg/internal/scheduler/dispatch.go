package scheduler

import (
	"fmt"
	"git.solsynth.dev/hypernet/nexus/pkg/nex/rx"
	"git.solsynth.dev/hypernet/pusher/pkg/internal/gap"
	"git.solsynth.dev/hypernet/pusher/pkg/internal/provider"
	"git.solsynth.dev/hypernet/pusher/pkg/pushkit"
	"github.com/go-playground/validator/v10"
	"github.com/goccy/go-json"
	"github.com/nats-io/nats.go"
)

var validate = validator.New(validator.WithRequiredStructEnabled())

func SubscribeToQueue() error {
	mq, err := rx.NewMqConn(gap.Nx)
	if err != nil {
		return fmt.Errorf("failed to initialize Nex.Rx connection: %v", err)
	}

	_, err = mq.Nt.Subscribe(pushkit.PushNotificationMqTopic, func(msg *nats.Msg) {
		var req pushkit.NotificationPushRequest
		if json.Unmarshal(msg.Data, &req) != nil {
			return
		} else if validate.Struct(&req) != nil {
			return
		}

		go provider.PushNotification(req)
	})
	if err != nil {
		return fmt.Errorf("failed to subscribe notification topic: %v", err)
	}

	_, err = mq.Nt.Subscribe(pushkit.PushNotificationBatchMqTopic, func(msg *nats.Msg) {
		var req pushkit.NotificationPushBatchRequest
		if json.Unmarshal(msg.Data, &req) != nil {
			return
		} else if validate.Struct(&req) != nil {
			return
		}

		go provider.PushNotificationBatch(req)
	})
	if err != nil {
		return fmt.Errorf("failed to subscribe notification batch topic: %v", err)
	}

	return nil
}
