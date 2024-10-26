package grpc

import (
	"context"
	"git.solsynth.dev/hypernet/pusher/pkg/internal/provider"
	"git.solsynth.dev/hypernet/pusher/pkg/proto"
	"git.solsynth.dev/hypernet/pusher/pkg/pushkit"
)

func (v *Server) PushNotification(ctx context.Context, request *proto.PushNotificationRequest) (*proto.DeliveryResponse, error) {
	err := provider.PushNotification(pushkit.NotificationPushRequest{
		Provider:     request.GetProvider(),
		Token:        request.GetDeviceToken(),
		Notification: pushkit.NewNotificationFromProto(request.GetNotify()),
	})
	return &proto.DeliveryResponse{IsSuccess: err == nil}, nil
}

func (v *Server) PushNotificationBatch(ctx context.Context, request *proto.PushNotificationBatchRequest) (*proto.DeliveryResponse, error) {
	go provider.PushNotificationBatch(pushkit.NotificationPushBatchRequest{
		Providers:    request.GetProviders(),
		Tokens:       request.GetDeviceTokens(),
		Notification: pushkit.NewNotificationFromProto(request.GetNotify()),
	})
	return &proto.DeliveryResponse{IsSuccess: true}, nil
}

func (v *Server) DeliverEmail(ctx context.Context, request *proto.DeliverEmailRequest) (*proto.DeliveryResponse, error) {
	go provider.SendMail(request.GetTo(), pushkit.NewEmailDataFromProto(request.GetEmail()))
	return &proto.DeliveryResponse{IsSuccess: true}, nil
}

func (v *Server) DeliverEmailBatch(ctx context.Context, request *proto.DeliverEmailBatchRequest) (*proto.DeliveryResponse, error) {
	go func() {
		for _, to := range request.GetTo() {
			_ = provider.SendMail(to, pushkit.NewEmailDataFromProto(request.GetEmail()))
		}
	}()
	return &proto.DeliveryResponse{IsSuccess: true}, nil
}
