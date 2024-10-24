package provider

import (
	"git.solsynth.dev/hypernet/pusher/pkg/pushkit"
)

type NotificationProvider interface {
	Push(in pushkit.Notification, tk string) error

	GetName() string
}
