package provider

import (
	"context"
	firebase "firebase.google.com/go"
	"firebase.google.com/go/messaging"
	"fmt"
	"git.solsynth.dev/hypernet/pusher/pkg/pushkit"
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

	var subtitle string
	if len(in.Subtitle) > 0 {
		subtitle = "\n" + in.Subtitle
	}
	message := &messaging.Message{
		Notification: &messaging.Notification{
			Title: in.Title,
			Body:  subtitle + in.Body,
		},
		Token: tk,
	}

	_, err = client.Send(ctx, message)
	return err
}

func (v *FirebaseNotifyProvider) GetName() string {
	return "fcm"
}
