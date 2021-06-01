package handler

type AppHandler interface {
	PostHandler
	CollectionHandler
	CommentHandler
	FavoriteHandler
	LikeHandler
	ReportPostHandler
}
