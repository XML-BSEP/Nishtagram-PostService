package usecase

import (
	"context"
	logger "github.com/jelena-vlajkov/logger/logger"
	"post-service/domain"
	"post-service/dto"
	"post-service/infrastructure/grpc/service/notification_service"
	pb "post-service/infrastructure/grpc/service/notification_service"
	"post-service/repository"
)

type LikeUseCase interface {
	LikePost(dto dto.LikeDislikeDTO, ctx context.Context) error
	DislikePost(dto dto.LikeDislikeDTO, ctx context.Context) error
	RemoveLike(dto dto.LikeDislikeDTO, ctx context.Context) error
	RemoveDislike(dto dto.LikeDislikeDTO, ctx context.Context) error
	GetLikesForPost(postId string, ctx context.Context) ([]dto.LikeDislikePreviewDTO, error)
	GetDislikesForPost(postId string, ctx context.Context) ([]dto.LikeDislikePreviewDTO, error)
}

type likeUseCase struct {
	likeRepository repository.LikeRepo
	logger *logger.Logger
	notificationClient notification_service.NotificationClient
}

func (l likeUseCase) RemoveLike(dto dto.LikeDislikeDTO, ctx context.Context) error {
	l.logger.Logger.Infof("removing like for user %v on post %v\n", dto.UserId, dto.PostId)

	err := l.likeRepository.RemoveLike(dto.PostId, dto.PostBy, domain.Profile{Id: dto.UserId}, context.Background())

	if err != nil {
		l.logger.Logger.Errorf("error while removing like for user %v on post %v, error: %v\n", dto.UserId, dto.PostId, err)
	}
	return err
}

func (l likeUseCase) LikePost(dto dto.LikeDislikeDTO, ctx context.Context) error {
	l.logger.Logger.Infof("adding like for user %v on post %v\n", dto.UserId, dto.PostId)
	err := l.likeRepository.LikePost(dto.PostId, dto.PostBy, domain.Profile{Id: dto.UserId}, context.Background())
	if err != nil {
		return err
	}

	if l.likeRepository.SeeIfDislikeExists(dto.PostId, dto.UserId, context.Background()) {
		l.logger.Logger.Infof("removing dislike for user %v on post %v\n", dto.UserId, dto.PostId)
		err = l.likeRepository.RemoveDislike(dto.PostId, dto.PostBy, domain.Profile{Id: dto.UserId}, context.Background())

		if err != nil {
			return err
		}
	}

	notification := &pb.NotificationMessage{
		Sender: dto.UserId,
		Receiver: dto.PostBy,
		NotificationType: pb.NotificationType_Like,
		RedirectPath: dto.PostId,
	}
	_, _ = l.notificationClient.SendNotification(ctx, notification)
	return nil
}

func (l likeUseCase) DislikePost(dto dto.LikeDislikeDTO, ctx context.Context) error {
	l.logger.Logger.Infof("adding dislike for user %v on post %v\n", dto.UserId, dto.PostBy)
	err := l.likeRepository.DislikePost(dto.PostId, dto.PostBy, domain.Profile{Id: dto.UserId}, context.Background())
	if err != nil {
		return err
	}

	if l.likeRepository.SeeIfLikeExists(dto.PostId, dto.UserId, context.Background()) {
		l.logger.Logger.Infof("removing like for user %v on post %v", dto.UserId, dto.PostId)
		err = l.likeRepository.RemoveLike(dto.PostId, dto.PostBy, domain.Profile{Id: dto.UserId}, context.Background())
		if err != nil {
			return err
		}
	}

	notification := &pb.NotificationMessage{
		Sender: dto.UserId,
		Receiver: dto.PostBy,
		NotificationType: pb.NotificationType_Dislike,
		RedirectPath: dto.PostId,
	}
	_, _ = l.notificationClient.SendNotification(ctx, notification)
	return nil
}

func (l likeUseCase) RemoveDislike(dto dto.LikeDislikeDTO, ctx context.Context) error {
	l.logger.Logger.Infof("removing dislike for user %v on post %v", dto.UserId, dto.PostId)
	return l.likeRepository.RemoveDislike(dto.PostId, dto.PostBy, domain.Profile{Id: dto.UserId}, context.Background())
}

func (l likeUseCase) GetLikesForPost(postId string, ctx context.Context) ([]dto.LikeDislikePreviewDTO, error) {
	l.logger.Logger.Infof("getting like for post %v\n", postId)
	likes, err := l.likeRepository.GetLikesForPost(postId, context.Background())
	if err != nil {
		l.logger.Logger.Errorf("error while getting likes for post %v, error: %v\n", postId, err)
		return nil, err
	}
	var dislikesPreview []dto.LikeDislikePreviewDTO
	for _, d := range likes {
		dislikesPreview = append(dislikesPreview, dto.NewLikeDislikePreviewDTO(postId, dto.NewUserPreviewDTO(d.Profile.Id)))
	}
	return dislikesPreview, nil
}

func (l likeUseCase) GetDislikesForPost(postId string, ctx context.Context) ([]dto.LikeDislikePreviewDTO, error) {
	l.logger.Logger.Infof("getting dislikes for post %v\n", postId)
	dislikes, err := l.likeRepository.GetDislikesForPost(postId, context.Background())
	if err != nil {
		l.logger.Logger.Errorf("error while getting dislikes for post %v, error: %v\n", postId, err)
		return nil, err
	}
	var dislikesPreview []dto.LikeDislikePreviewDTO
	for _, d := range dislikes {
		dislikesPreview = append(dislikesPreview, dto.NewLikeDislikePreviewDTO(postId, dto.NewUserPreviewDTO(d.Profile.Id)))
	}
	return dislikesPreview, nil
}


func NewLikeUseCase(likeRepository repository.LikeRepo, logger *logger.Logger, notificationClient notification_service.NotificationClient) LikeUseCase {
	return &likeUseCase{likeRepository: likeRepository, logger: logger, notificationClient: notificationClient}
}