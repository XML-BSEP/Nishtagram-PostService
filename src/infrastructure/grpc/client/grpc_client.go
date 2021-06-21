package client

import (
	"google.golang.org/grpc"
	pb "post-service/infrastructure/grpc/service/notification_service"
)

func NewNotificationClient(address string) (pb.NotificationClient, error) {

	conn, err := grpc.Dial(address, grpc.WithInsecure())

	if err != nil {
		return nil, err
	}

	client := pb.NewNotificationClient(conn)
	return client, nil
}

