package provider

import (
	"fmt"
	"sync"
	"time"

	"git.solsynth.dev/hypernet/pusher/pkg/pushkit"
	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
)

var notifyProviders = make(map[string]NotificationProvider)

func AddProvider(in NotificationProvider) {
	notifyProviders[in.GetName()] = in
}

func PushNotification(in pushkit.NotificationPushRequest) error {
	in.Notification = TranslateNotify(in.Notification, in.Lang)

	requestId := uuid.NewString()
	log.Debug().
		Str("tk", in.Token).
		Str("provider", in.Provider).
		Str("topic", in.Notification.Topic).
		Str("request_id", requestId).
		Any("body", in.Notification).
		Msg("Pushing notification...")

	prov, ok := notifyProviders[in.Provider]
	if !ok {
		log.Warn().
			Str("provider", in.Provider).
			Str("topic", in.Notification.Topic).
			Str("request_id", requestId).
			Msg("Provider was not found, push skipped...")
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
		Any("tk", in.Tokens).
		Any("providers", in.Providers).
		Str("topic", in.Notification.Topic).
		Any("body", in.Notification).
		Str("request_id", requestId).
		Msg("Pushing notification in batch...")

	var wg sync.WaitGroup
	for idx, key := range in.Providers {
		prov, ok := notifyProviders[key]
		if !ok {
			log.Warn().
				Str("provider", key).
				Str("topic", in.Notification.Topic).
				Str("request_id", requestId).
				Msg("Provider was not found, push skipped...")
			continue
		}
		wg.Add(1)
		go func() {
			defer wg.Done()

			in.Notification = TranslateNotify(in.Notification, in.Lang[idx])
			log.Debug().
				Str("tk", in.Tokens[0]).
				Str("provider", in.Providers[0]).
				Str("topic", in.Notification.Topic).
				Str("request_id", requestId).
				Msg("Pushing notification...")

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
