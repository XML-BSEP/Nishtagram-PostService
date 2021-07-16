package interactor

import (
	"github.com/gocql/gocql"
	logger "github.com/jelena-vlajkov/logger/logger"
	"post-service/http/handler"
	"post-service/infrastructure/data_seeder"
	"post-service/infrastructure/grpc/service/notification_service"
	"post-service/repository"
	"post-service/usecase"
)

type Interactor interface {
	NewPostRepo() repository.PostRepo
	NewLikeRepo() repository.LikeRepo
	NewFavoriteRepo() repository.FavoritesRepo
	NewCollectionRepo() repository.CollectionRepo
	NewCommentRepo() repository.CommentRepo
	NewReportRepo() repository.ReportRepo
	NewMutedRepo() repository.MutedContentRepo

	NewPostUseCase() usecase.PostUseCase
	NewReportPostUseCase() usecase.PostReportUseCase
	NewLikeUseCase() usecase.LikeUseCase
	NewCommentUseCase() usecase.CommentUseCase
	NewFavoriteUseCase() usecase.FavoriteUseCase
	NewCollectionUseCase() usecase.CollectionUseCase
	NewMutedUseCase() usecase.MutedContentUseCase

	NewAppHandler() handler.AppHandler
	NewPostHandler() handler.PostHandler
	NewReportPostHandler() handler.ReportPostHandler
	NewLikeHandler() handler.LikeHandler
	NewCommentHandler() handler.CommentHandler
	NewFavoriteHandler() handler.FavoriteHandler
	NewCollectionHandler() handler.CollectionHandler
	NewMutedHandler() handler.MutedContentHandler


}

type interactor struct {
	cassandraSession *gocql.Session
	logger *logger.Logger
	notificationClient notification_service.NotificationClient
}

func (i interactor) NewMutedRepo() repository.MutedContentRepo {
	return repository.NewMutedContentRepo(i.cassandraSession)
}

func (i interactor) NewMutedUseCase() usecase.MutedContentUseCase {
	return usecase.NewMutedContentUseCase(i.NewMutedRepo())
}

func (i interactor) NewMutedHandler() handler.MutedContentHandler {
	return handler.NewMuteContentHandler(i.NewMutedUseCase())
}

func (i interactor) NewPostUseCase() usecase.PostUseCase {
	return usecase.NewPostUseCase(i.NewPostRepo(), i.NewLikeRepo(), i.NewFavoriteRepo(), i.NewCollectionRepo(), i.logger, i.notificationClient, i.NewMutedUseCase())
}

func (i interactor) NewReportPostUseCase() usecase.PostReportUseCase {
	return usecase.NewPostReportUseCase(i.NewReportRepo())
}

func (i interactor) NewLikeUseCase() usecase.LikeUseCase {
	return usecase.NewLikeUseCase(i.NewLikeRepo(), i.logger, i.notificationClient)
}

func (i interactor) NewCommentUseCase() usecase.CommentUseCase {
	return usecase.NewCommentUseCase(i.NewCommentRepo(), i.logger, i.notificationClient)
}

func (i interactor) NewFavoriteUseCase() usecase.FavoriteUseCase {
	return usecase.NewFavoriteUseCase(i.NewFavoriteRepo(), i.NewPostRepo(), i.NewPostUseCase(), i.logger)
}

func (i interactor) NewCollectionUseCase() usecase.CollectionUseCase {
	return usecase.NewCollectionUseCase(i.NewCollectionRepo(), i.NewPostRepo(), i.NewPostUseCase(), i.logger)
}

func (i interactor) NewReportPostHandler() handler.ReportPostHandler {
	return handler.NewReportPostHandler(i.NewReportPostUseCase(), i.logger)
}

func (i interactor) NewLikeHandler() handler.LikeHandler {
	return handler.NewLikeHandler(i.NewLikeUseCase(), i.logger)
}

func (i interactor) NewCommentHandler() handler.CommentHandler {
	return handler.NewCommentHandler(i.NewCommentUseCase(), i.logger)
}

func (i interactor) NewFavoriteHandler() handler.FavoriteHandler {
	return handler.NewFavoriteHandler(i.NewFavoriteUseCase(), i.logger)
}

func (i interactor) NewCollectionHandler() handler.CollectionHandler {
	return handler.NewCollectionHandler(i.NewCollectionUseCase(), i.logger)
}

type appHandler struct {
	handler.PostHandler
	handler.LikeHandler
	handler.CommentHandler
	handler.FavoriteHandler
	handler.ReportPostHandler
	handler.CollectionHandler
	handler.MutedContentHandler
}

func (i interactor) NewAppHandler() handler.AppHandler {
	appHandler := appHandler{}
	appHandler.PostHandler = i.NewPostHandler()
	appHandler.LikeHandler = i.NewLikeHandler()
	appHandler.ReportPostHandler = i.NewReportPostHandler()
	appHandler.CommentHandler = i.NewCommentHandler()
	appHandler.CollectionHandler = i.NewCollectionHandler()
	appHandler.FavoriteHandler = i.NewFavoriteHandler()
	appHandler.MutedContentHandler = i.NewMutedHandler()

	data_seeder.SeedData(i.cassandraSession)

	return appHandler
}

func (i interactor) NewPostHandler() handler.PostHandler {
	return handler.NewPostHandler(i.NewPostUseCase(), i.logger)
}


func (i interactor) NewPostRepo() repository.PostRepo {
	return repository.NewPostRepository(i.cassandraSession, i.logger)
}

func (i interactor) NewLikeRepo() repository.LikeRepo {
	return repository.NewLikeRepository(i.cassandraSession, i.logger)
}

func (i interactor) NewFavoriteRepo() repository.FavoritesRepo {
	return repository.NewFavoritesRepository(i.cassandraSession, i.logger)
}

func (i interactor) NewCollectionRepo() repository.CollectionRepo {
	return repository.NewCollectionRepository(i.cassandraSession, i.logger)
}

func (i interactor) NewCommentRepo() repository.CommentRepo {
	return repository.NewCommentRepository(i.cassandraSession, i.logger)
}

func (i interactor) NewReportRepo() repository.ReportRepo {
	return repository.NewReportRepository(i.cassandraSession)
}

func NewInteractor(cassandraSession *gocql.Session, logger *logger.Logger, notificationClient notification_service.NotificationClient) Interactor {
	return &interactor{cassandraSession: cassandraSession, logger: logger, notificationClient: notificationClient}
}
