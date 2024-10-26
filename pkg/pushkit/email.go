package pushkit

import (
	"git.solsynth.dev/hypernet/pusher/pkg/proto"
	"github.com/samber/lo"
)

type EmailDeliverRequest struct {
	To    string    `json:"to" validate:"required"`
	Email EmailData `json:"email" validate:"required"`
}

type EmailDeliverBatchRequest struct {
	To    []string  `json:"to" validate:"required"`
	Email EmailData `json:"email" validate:"required"`
}

type EmailData struct {
	Subject string  `json:"subject" validate:"required"`
	Text    *string `json:"text"`
	HTML    *string `json:"html"`
}

func NewEmailDataFromProto(in *proto.EmailInfo) EmailData {
	return EmailData{
		Subject: in.GetSubject(),
		Text:    lo.ToPtr(in.GetTextBody()),
		HTML:    lo.ToPtr(in.GetHtmlBody()),
	}
}
