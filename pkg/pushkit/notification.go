package pushkit

import (
	"git.solsynth.dev/hypernet/nexus/pkg/nex"
	"git.solsynth.dev/hypernet/pusher/pkg/proto"
)

type NotificationPushRequest struct {
	Provider     string       `json:"provider" validate:"required"`
	Token        string       `json:"token" validate:"required"`
	Notification Notification `json:"notification" validate:"required"`
}

type NotificationPushBatchRequest struct {
	Providers    []string     `json:"provider" validate:"required"`
	Tokens       []string     `json:"tokens" validate:"required"`
	Notification Notification `json:"notification" validate:"required"`
}

type Notification struct {
	Topic    string         `json:"topic" validate:"required"`
	Title    string         `json:"title" validate:"required"`
	Subtitle string         `json:"subtitle"`
	Body     string         `json:"body" validate:"required"`
	Metadata map[string]any `json:"metadata"`
	Priority int            `json:"priority"`
}

func NewNotificationFromProto(in *proto.NotifyInfo) Notification {
	return Notification{
		Topic:    in.GetTopic(),
		Title:    in.GetTitle(),
		Subtitle: in.GetSubtitle(),
		Body:     in.GetBody(),
		Metadata: nex.DecodeMap(in.GetMetadata()),
		Priority: int(in.GetPriority()),
	}
}
