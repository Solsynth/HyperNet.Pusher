package provider

import (
	"git.solsynth.dev/hypernet/pusher/pkg/pushkit"
	"github.com/sideshow/apns2"
	payload2 "github.com/sideshow/apns2/payload"
	"github.com/spf13/viper"
)

type AppleNotifyProvider struct {
	topic string
	conn  *apns2.Client
}

func (v *AppleNotifyProvider) Push(in pushkit.Notification, tk string) error {
	data := payload2.
		NewPayload().
		AlertTitle(in.Title).
		AlertBody(in.Body).
		Category(in.Topic).
		Custom("metadata", in.Metadata).
		Sound("default").
		MutableContent()
	if len(in.Subtitle) > 0 {
		data = data.AlertSubtitle(in.Subtitle)
	}
	if avatar, ok := in.Metadata["avatar"]; ok {
		data = data.Custom("avatar", avatar)
	}
	if picture, ok := in.Metadata["picture"]; ok {
		data = data.Custom("picture", picture)
	}
	rawData, err := data.MarshalJSON()
	if err != nil {
		return err
	}
	payload := &apns2.Notification{
		DeviceToken: tk,
		Topic:       viper.GetString(v.topic),
		Payload:     rawData,
	}
	_, err = v.conn.Push(payload)
	return err
}

func (v *AppleNotifyProvider) GetName() string {
	return "apns"
}
