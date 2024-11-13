package provider

import (
	"fmt"
	"git.solsynth.dev/hypernet/pusher/pkg/pushkit"
	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
	"sync"
	"time"
)

var notifyProviders = make(map[string]NotificationProvider)

func AddProvider(in NotificationProvider) {
	notifyProviders[in.GetName()] = in
}

func PushNotification(in pushkit.NotificationPushRequest) error {
	requestId := uuid.NewString()
	log.Debug().
		Str("tk", in.Token).
		Str("provider", in.Provider).
		Str("topic", in.Notification.Topic).
		Str("request_id", requestId).
		Msg("Pushing notification...")

	prov, ok := notifyProviders[in.Provider]
	if !ok {
		return fmt.Errorf("provider not found")
	}
	start := time.Now()
	err := prov.Push(in.Notification, in.Token)
	if err != nil {
		log.Warn().Err(err).
			Str("tk", in.Token).
			Str("provider", prov.GetName()).
			Dur("elapsed", time.Since(start)).
			Str("request_id", requestId).
			Msg("Push notification failed once")
	} else {
		log.Debug().
			Str("tk", in.Token).
			Str("provider", prov.GetName()).
			Dur("elapsed", time.Since(start)).
			Str("request_id", requestId).
			Msg("Pushed one notification")
	}
	return err
}

func PushNotificationBatch(in pushkit.NotificationPushBatchRequest) {
	requestId := uuid.NewString()
	log.Debug().
		Int("count", len(in.Tokens)).
		Str("topic", in.Notification.Topic).
		Str("request_id", requestId).
		Msg("Pushing notification batch...")

	var wg sync.WaitGroup
	for idx, key := range in.Providers {
		prov, ok := notifyProviders[key]
		if !ok {
			continue
		}
		go func() {
			wg.Add(1)
			defer wg.Done()
			start := time.Now()
			err := prov.Push(in.Notification, in.Tokens[idx])
			if err != nil {
				log.Warn().Err(err).
					Str("tk", in.Tokens[idx]).
					Str("provider", prov.GetName()).
					Dur("elapsed", time.Since(start)).
					Str("request_id", requestId).
					Msg("Push notification failed once")
			} else {
				log.Debug().
					Str("tk", in.Tokens[idx]).
					Str("provider", prov.GetName()).
					Dur("elapsed", time.Since(start)).
					Str("request_id", requestId).
					Msg("Pushed one notification")
			}
		}()
	}
}
