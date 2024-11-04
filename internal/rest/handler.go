package rest

import (
	"net/http"

	"github.com/Gen1usBruh/MiniTwitter/internal/middleware"
	"github.com/Gen1usBruh/MiniTwitter/internal/scope"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "github.com/Gen1usBruh/MiniTwitter/docs"
)

type HandlerConfig struct {
	Dep *scope.Dependencies
}

func NewHandler(cfg HandlerConfig) http.Handler {
	router := gin.Default()

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.Use(middleware.AuthMiddleware())

	timeline := router.Group("/timeline")
	{
		timeline.GET("/", cfg.getTimeline)
	}

	likes := router.Group("/likes")
	{
		likes.POST("/tweet", cfg.createLikeTweet)
		likes.DELETE("/tweet", cfg.deleteLikeTweet)
		likes.GET("/tweet", cfg.getLikeTweet)

		likes.POST("/comment/:comment_id", cfg.createLikeComment)
		likes.DELETE("/comment/:comment_id", cfg.deleteLikeComment)
		likes.GET("/comment/:comment_id", cfg.getLikeComment)
	}

	comments := router.Group("/comment")
	{
		comments.POST("/", cfg.createComment)
		comments.GET("/", cfg.getComment)
		comments.DELETE("/:comment_id", cfg.deleteComment)
	}

	users := router.Group("/users")
	{
		users.POST("/signup", cfg.signUp)
		users.POST("/login", cfg.signIn)
		users.GET("/:user_id", cfg.getUser)
		users.DELETE("/:user_id", cfg.deleteUser)
		users.PUT("/:user_id", cfg.updateUser)
		users.GET("/:user_id/followers", cfg.getUserFollowers)
		users.GET("/:user_id/following", cfg.getUserFollowing)
	}

	tweets := router.Group("/tweets")
	{
		tweets.POST("/", cfg.createTweet)
		tweets.GET("/:tweet_id", cfg.getTweet)
		tweets.DELETE("/:tweet_id", cfg.deleteTweet)
	}

	retweets := router.Group("/retweets")
	{
		retweets.POST("/", cfg.createRetweet)
		retweets.GET("/:retweet_id", cfg.getRetweet)
		retweets.DELETE("/:retweet_id", cfg.deleteRetweet)
		retweets.GET("/tweet/:id", cfg.getRetweetTweet)
		retweets.GET("/retweet/:id", cfg.getRetweetRetweet)
	}

	follow := router.Group("/follow")
	{
		follow.POST("/", cfg.follow)
		follow.DELETE("/", cfg.unfollow)
	}

	return router
}