package provider

import (
	"context"
	firebase "firebase.google.com/go"
	"github.com/sideshow/apns2"
	"github.com/sideshow/apns2/token"
	"google.golang.org/api/option"
)

func InitFCM(in string) error {
	opt := option.WithCredentialsFile(in)
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		return err
	} else {
		AddProvider(&FirebaseNotifyProvider{app})
	}

	return nil
}

func InitAPN(in, keyId, teamId, topic string) error {
	authKey, err := token.AuthKeyFromFile(in)
	if err != nil {
		return err
	} else {
		AddProvider(&AppleNotifyProvider{topic, apns2.NewTokenClient(&token.Token{
			AuthKey: authKey,
			KeyID:   keyId,
			TeamID:  teamId,
		}).Production()})
	}

	return nil
}
