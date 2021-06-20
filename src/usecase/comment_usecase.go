package usecase

import (
	"context"
	logger "github.com/jelena-vlajkov/logger/logger"
	"post-service/dto"
	"post-service/infrastructure/grpc/service/notification_service"
	"post-service/repository"
	pb "post-service/infrastructure/grpc/service/notification_service"
)

type CommentUseCase interface {
	AddComment(comment dto.CommentDTO, ctx context.Context) error
	DeleteComment(comment dto.CommentDTO, ctx context.Context) error
	GetAllCommentsByPost(postId string, ctx context.Context) ([]dto.CommentDTO, error)
}

type commentUseCase struct {
	commentRepository repository.CommentRepo
	logger *logger.Logger
	notificationClient notification_service.NotificationClient
}

func (c commentUseCase) GetAllCommentsByPost(postId string, ctx context.Context) ([]dto.CommentDTO, error) {
	c.logger.Logger.Infof("getting comments for post %v\n", postId)
	comments, err := c.commentRepository.GetComments(postId, context.Background())
	if err != nil {
		c.logger.Logger.Errorf("error while getting comments for post %v, error: %v\n", postId, err)
	}

	return comments, err
}

func (c commentUseCase) AddComment(comment dto.CommentDTO, ctx context.Context) error {
	c.logger.Logger.Infof("adding comment on post %v by user %v\n", comment.PostId, comment.PostBy)
	err := c.commentRepository.CommentPost(comment, context.Background())
	notification := &pb.NotificationMessage{
		Sender: comment.CommentBy.Id,
		Receiver: comment.PostBy,
		NotificationType: pb.NotificationType_Like,
		RedirectPath: comment.PostId,
	}
	_, _ = c.notificationClient.SendNotification(ctx, notification)
	return err
}

func (c commentUseCase) DeleteComment(comment dto.CommentDTO, ctx context.Context) error {
	c.logger.Logger.Infof("deleting comment on post %v with id %v\n", comment.PostId, comment.CommentId)
	err := c.commentRepository.DeleteComment(comment, context.Background())
	if err != nil {
		c.logger.Logger.Errorf("error while deleting comment on post %v with id %v, error: %v\n", comment.PostId, comment.PostId, err)
	}
	return err
}

func NewCommentUseCase(commentRepository repository.CommentRepo, logger *logger.Logger, client notification_service.NotificationClient) CommentUseCase {
	return &commentUseCase{
		commentRepository: commentRepository,
		logger: logger,
		notificationClient: client,
	}
}