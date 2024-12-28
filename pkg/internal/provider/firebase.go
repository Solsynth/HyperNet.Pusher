package provider

import (
	"context"
	"fmt"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/messaging"
	"git.solsynth.dev/hypernet/pusher/pkg/pushkit"
	"github.com/rs/zerolog/log"
)

type FirebaseNotifyProvider struct {
	conn *firebase.App
}

func (v *FirebaseNotifyProvider) Push(in pushkit.Notification, tk string) error {
	ctx := context.Background()
	client, err := v.conn.Messaging(ctx)
	if err != nil {
		return fmt.Errorf("failed to create firebase client")
	}

	var body string
	if len(in.Subtitle) > 0 {
		body = fmt.Sprintf("%s\n%s", in.Body, in.Subtitle)
	} else {
		body = in.Body
	}

	message := &messaging.Message{
		Notification: &messaging.Notification{
			Title: in.Title,
			Body:  body,
		},
		Token: tk,
	}

	resp, err := client.Send(ctx, message)
	log.Debug().
		Str("token", tk).
		Str("response", resp).
		Msg("Pushed once notification to firebase")

	return err
}

func (v *FirebaseNotifyProvider) GetName() string {
	return "fcm"
}
