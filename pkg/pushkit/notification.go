package pushkit

import (
	"git.solsynth.dev/hypernet/nexus/pkg/nex"
	"git.solsynth.dev/hypernet/pusher/pkg/proto"
	"github.com/goccy/go-json"
)

type NotificationPushRequest struct {
	Lang         string       `json:"language" validate:"required"`
	Provider     string       `json:"provider" validate:"required"`
	Token        string       `json:"token" validate:"required"`
	Notification Notification `json:"notification" validate:"required"`
}

type NotificationPushBatchRequest struct {
	Lang         []string     `json:"language" validate:"required"`
	Providers    []string     `json:"provider" validate:"required"`
	Tokens       []string     `json:"tokens" validate:"required"`
	Notification Notification `json:"notification" validate:"required"`
}

type Notification struct {
	TranslateKey  *string             `json:"tr_key"`
	TranslateArgs map[string][]string `json:"tr_args"`
	Topic         string              `json:"topic" validate:"required"`
	Title         string              `json:"title" validate:"required"`
	Subtitle      string              `json:"subtitle"`
	Body          string              `json:"body" validate:"required"`
	Metadata      map[string]any      `json:"metadata"`
	Priority      int                 `json:"priority"`
}

func NewNotificationFromProto(in *proto.NotifyInfo) Notification {
	var args map[string][]string
	_ = json.Unmarshal(in.GetTranslateArgs(), &args)

	return Notification{
		Topic:         in.GetTopic(),
		Title:         in.GetTitle(),
		Subtitle:      in.GetSubtitle(),
		Body:          in.GetBody(),
		Metadata:      nex.DecodeMap(in.GetMetadata()),
		Priority:      int(in.GetPriority()),
		TranslateKey:  in.TranslateKey,
		TranslateArgs: args,
	}
}
