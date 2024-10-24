package provider

import (
	"fmt"
	"git.solsynth.dev/hypernet/pusher/pkg/pushkit"
	"github.com/rs/zerolog/log"
	"sync"
	"time"
)

var notifyProviders = make(map[string]NotificationProvider)

func AddProvider(in NotificationProvider) {
	notifyProviders[in.GetName()] = in
}

func PushNotification(in pushkit.NotificationPushRequest) error {
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
			Msg("Push notification failed once")
	} else {
		log.Debug().
			Str("tk", in.Token).
			Str("provider", prov.GetName()).
			Dur("elapsed", time.Since(start)).
			Msg("Pushed one notification")
	}
	return err
}

func PushNotificationBatch(in pushkit.NotificationPushBatchRequest) {
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
					Msg("Push notification failed once")
			} else {
				log.Debug().
					Str("tk", in.Tokens[idx]).
					Str("provider", prov.GetName()).
					Dur("elapsed", time.Since(start)).
					Msg("Pushed one notification")
			}
		}()
	}
}
