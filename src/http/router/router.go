package router

import (
	"github.com/gin-gonic/gin"
	logger "github.com/jelena-vlajkov/logger/logger"
	"post-service/http/handler"
	"post-service/http/middleware"
	"post-service/http/middleware/prometheus_middleware"
)

func NewRouter(handler handler.AppHandler, logger *logger.Logger) *gin.Engine{
	router := gin.Default()
	requestCounter := prometheus_middleware.GetHttpRequestsCounter()
	router.Use(prometheus_middleware.PrometheusMiddleware(requestCounter))
	router.GET("/metrics", prometheus_middleware.PrometheusGinHandler())

	g := router.Group("/post")

	g.Use(middleware.AuthMiddleware(logger))

	g.POST("createPost", handler.AddPost)
	g.POST("createCollection", handler.CreateCollection)
	g.POST("reportPost", handler.ReportPost)
	g.POST("likePost", handler.LikePost)
	g.POST("dislikePost", handler.DislikePost)
	g.POST("removeLike", handler.RemoveLike)
	g.POST("removeDislike", handler.RemoveDislike)
	g.POST("comment", handler.AddComment)
	g.POST("removeComment", handler.DeleteComment)
	g.POST("addToCollection", handler.AddPostToCollection)
	g.POST("removeFromCollection", handler.RemovePostFromCollection)
	g.POST("addFavorite", handler.AddPostToFavorite)
	g.POST("removeFavorite", handler.RemovePostFromFavorites)
	g.POST("deleteCollection", handler.DeleteCollection)
	g.POST("reviewReport", handler.ReviewReport)
	g.POST("getPostsByUser", handler.GetPostsByUser)
	g.GET("getPendingReports", handler.GetAllPendingReports)
	g.GET("getApprovedReports", handler.GetAllApprovedReports)
	g.GET("getRejectedReports", handler.GetAllRejectedReports)
	g.POST("editPost", handler.EditPost)
	g.POST("deletePost", handler.DeletePost)
	g.GET("feed", handler.GenerateUserFeed)
	g.POST("getAllComments", handler.GetComments)
	g.POST("getPostsOnProfile", handler.GetPostsOnProfile)
	g.POST("getPostById", handler.GetPostById)
	g.GET("getAllCollections", handler.GetAllCollections)
	g.POST("getPostsInCollection", handler.GetCollection)
	g.GET("getFavorites", handler.GetFavorites)
	g.GET("likedMedia", handler.GetLikedMedia)
	g.GET("dislikedMedia", handler.GetDislikedMedia)
	g.POST("getPostByIdForSearch", handler.GetPostByIdForSearch)
	g.POST("mute", handler.MuteUser)
	g.POST("unmute", handler.UnmuteUser)
	g.POST("isMuted", handler.SeeIfMuted)
	g.POST("createPostFromCampaign", handler.AddPostFromCampaign)

	g.GET("getAllReportTypes", handler.GetAllReportTypes)


	return router
}