package client

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	pb "post-service/infrastructure/grpc/service/notification_service"
)

func NewNotificationClient(address string) (pb.NotificationClient, error) {
	creds, err := credentials.NewClientTLSFromFile("certificate/cert.pem", "")
	if err != nil {
		return nil, err
	}
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(creds))

	if err != nil {
		return nil, err
	}

	client := pb.NewNotificationClient(conn)
	return client, nil
}

