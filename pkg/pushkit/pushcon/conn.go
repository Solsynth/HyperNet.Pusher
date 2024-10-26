package pushcon

import (
	"git.solsynth.dev/hypernet/nexus/pkg/nex"
	"git.solsynth.dev/hypernet/nexus/pkg/nex/rx"
	"git.solsynth.dev/hypernet/pusher/pkg/pushkit"
)

type Conn struct {
	n  *nex.Conn
	mq *rx.MqConn
}

func NewConn(conn *nex.Conn) (*Conn, error) {
	c := &Conn{
		n: conn,
	}

	if mq, err := rx.NewMqConn(conn); err != nil {
		return nil, err
	} else {
		c.mq = mq
	}

	return c, nil
}

func (v *Conn) PushNotify(in pushkit.NotificationPushRequest) error {
	return v.mq.Nt.Publish(pushkit.PushNotificationMqTopic, nex.EncodeMap(in))
}

func (v *Conn) PushNotifyBatch(in pushkit.NotificationPushBatchRequest) error {
	return v.mq.Nt.Publish(pushkit.PushNotificationBatchMqTopic, nex.EncodeMap(in))
}

func (v *Conn) PushEmail(in pushkit.EmailDeliverRequest) error {
	return v.mq.Nt.Publish(pushkit.PushEmailMqTopic, nex.EncodeMap(in))
}

func (v *Conn) PushEmailBatch(in pushkit.EmailDeliverBatchRequest) error {
	return v.mq.Nt.Publish(pushkit.PushEmailBatchMqTopic, nex.EncodeMap(in))
}
