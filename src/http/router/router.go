package router

import (
	"github.com/gin-gonic/gin"
	"post-service/http/handler"
	"post-service/http/middleware"
)

func NewRouter(handler handler.AppHandler) *gin.Engine{
	router := gin.Default()

	g := router.Group("/post")

	g.Use(middleware.AuthMiddleware())

	g.POST("createPost", handler.AddPost)
	g.POST("createCollection", handler.CreateCollection)
	g.POST("createReport", handler.ReportPost)
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
	g.POST("getPendingReports", handler.GetAllPendingReports)
	g.POST("getApprovedReports", handler.GetAllApprovedReports)
	g.POST("getRejectedReports", handler.GetAllRejectedReports)
	g.POST("editPost", handler.EditPost)
	g.POST("deletePost", handler.DeletePost)
	g.GET("feed", handler.GenerateUserFeed)
	g.POST("getAllComments", handler.GetComments)
	g.POST("getPostsOnProfile", handler.GetPostsOnProfile)
	g.POST("getPostById", handler.GetPostById)
	g.GET("getAllCollections", handler.GetAllCollections)
	g.POST("getPostsInCollection", handler.GetCollection)
	g.GET("getFavorites", handler.GetFavorites)



	return router
}