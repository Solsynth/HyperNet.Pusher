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
	"github.com/rs/zerolog/log"
)

var validate = validator.New(validator.WithRequiredStructEnabled())

func SubscribeToQueue() error {
	mq, err := rx.NewMqConn(gap.Nx)
	if err != nil {
		return fmt.Errorf("failed to initialize Nex.Rx connection: %v", err)
	}

	_, err = mq.Nt.Subscribe(pushkit.PushNotificationMqTopic, func(msg *nats.Msg) {
		var req pushkit.NotificationPushRequest
		if err := json.Unmarshal(msg.Data, &req); err != nil {
			log.Warn().Err(err).Msg("Dropped a notify request, unable to parse request body")
			return
		} else if err := validate.Struct(&req); err != nil {
			log.Warn().Err(err).Msg("Dropped a notify request, failed to validate request body")
			return
		}

		go provider.PushNotification(req)
	})
	if err != nil {
		return fmt.Errorf("failed to subscribe notification topic: %v", err)
	}

	_, err = mq.Nt.Subscribe(pushkit.PushNotificationBatchMqTopic, func(msg *nats.Msg) {
		var req pushkit.NotificationPushBatchRequest
		if err := json.Unmarshal(msg.Data, &req); err != nil {
			log.Warn().Err(err).Msg("Dropped a notify batch request, unable to parse request body")
			return
		} else if err := validate.Struct(&req); err != nil {
			log.Warn().Err(err).Msg("Dropped a notify batch request, failed to validate request body")
			return
		}

		go provider.PushNotificationBatch(req)
	})
	if err != nil {
		return fmt.Errorf("failed to subscribe notification batch topic: %v", err)
	}

	_, err = mq.Nt.Subscribe(pushkit.PushEmailMqTopic, func(msg *nats.Msg) {
		var req pushkit.EmailDeliverRequest
		if err := json.Unmarshal(msg.Data, &req); err != nil {
			log.Warn().Err(err).Msg("Dropped a push email request, unable to parse request body")
			return
		} else if err := validate.Struct(&req); err != nil {
			log.Warn().Err(err).Msg("Dropped a push email request, failed to validate request body")
			return
		}

		go provider.SendMail(req.To, req.Email)
	})
	if err != nil {
		return fmt.Errorf("failed to subscribe email topic: %v", err)
	}

	_, err = mq.Nt.Subscribe(pushkit.PushEmailBatchMqTopic, func(msg *nats.Msg) {
		var req pushkit.EmailDeliverBatchRequest
		if err := json.Unmarshal(msg.Data, &req); err != nil {
			log.Warn().Err(err).Msg("Dropped a push email batch request, unable to parse request body")
			return
		} else if err := validate.Struct(&req); err != nil {
			log.Warn().Err(err).Msg("Dropped a push email batch request, failed to validate request body")
			return
		}

		go func() {
			for _, to := range req.To {
				_ = provider.SendMail(to, req.Email)
			}
		}()
	})
	if err != nil {
		return fmt.Errorf("failed to subscribe email batch topic: %v", err)
	}

	return nil
}
