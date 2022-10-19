package router

import (
	"final-project/controllers"
	"final-project/middlewares"
	"net/http"

	"github.com/gin-gonic/gin"
)

func StartApp() *gin.Engine {
	r := gin.Default()

	userRouter := r.Group("/users")
	{
		userRouter.POST("/register", controllers.CreateUser)
		userRouter.POST("/login", controllers.LoginUser)
		userRouter.DELETE("/", middlewares.Authentication(), controllers.DeleteUser)
		userRouter.PUT("/", middlewares.Authentication(), controllers.UpdateUser)
	}

	photoRouter := r.Group("/photos")
	{
		photoRouter.Use(middlewares.Authentication())
		photoRouter.POST("/", controllers.CreatePhoto)
		photoRouter.GET("/", controllers.GetAllPhotos)
		photoRouter.PUT("/:photoID", middlewares.PhotoAuthorization(), controllers.UpdatePhoto)
		photoRouter.DELETE("/:photoID", middlewares.PhotoAuthorization(), controllers.DeletePhoto)
	}

	commentRouter := r.Group("/comments")
	{
		commentRouter.Use(middlewares.Authentication())
		commentRouter.POST("/", controllers.CreateComment)
		commentRouter.GET("/", controllers.GetAllComments)
		commentRouter.PUT("/:commentID", middlewares.CommentAuthorization(), controllers.UpdateComment)
		commentRouter.DELETE("/:commentID", middlewares.CommentAuthorization(), controllers.DeleteComment)
	}

	socmedRouter := r.Group("/socialmedias")
	{
		socmedRouter.Use(middlewares.Authentication())
		socmedRouter.POST("/", controllers.CreateSocialMedia)
		socmedRouter.GET("/", controllers.GetAllSocialMedias)
		socmedRouter.PUT("/:socialMediaID", middlewares.SocialMediaAuthorization(), controllers.UpdateSocialMedia)
		socmedRouter.DELETE("/:socialMediaID", middlewares.SocialMediaAuthorization(), controllers.DeleteSocialMedia)

	}

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, "Server is running. Welcome to my OpenMyGram API")
	})

	return r
}
