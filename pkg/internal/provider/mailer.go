package provider

import (
	"crypto/tls"
	"fmt"
	"git.solsynth.dev/hypernet/pusher/pkg/pushkit"
	"github.com/jordan-wright/email"
	"github.com/spf13/viper"
	"net/smtp"
	"net/textproto"
)

func SendMail(target string, in pushkit.EmailData) error {
	mail := &email.Email{
		To:      []string{target},
		From:    viper.GetString("mailer.name"),
		Subject: in.Subject,
		Headers: textproto.MIMEHeader{},
	}
	if in.Text != nil {
		mail.Text = []byte(*in.Text)
	}
	if in.HTML != nil {
		mail.HTML = []byte(*in.HTML)
	}
	return mail.SendWithTLS(
		fmt.Sprintf("%s:%d", viper.GetString("mailer.smtp_host"), viper.GetInt("mailer.smtp_port")),
		smtp.PlainAuth(
			"",
			viper.GetString("mailer.username"),
			viper.GetString("mailer.password"),
			viper.GetString("mailer.smtp_host"),
		),
		&tls.Config{ServerName: viper.GetString("mailer.smtp_host")},
	)
}
