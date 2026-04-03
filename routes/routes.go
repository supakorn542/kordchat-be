package routes

import (
	"github.com/gin-gonic/gin"
	"kordchat-be/controllers"
	"kordchat-be/middlewares"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetupRoutes(r *gin.Engine) {

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	api := r.Group("api")
	{
		api.POST("/register", controllers.Register)
		api.POST("/login", controllers.Login)
		api.POST("/refresh", controllers.RefreshToken)

		protected := api.Group("/")
		protected.Use(middlewares.RequireAuth)
		{
			protected.POST("/servers", controllers.CreateServer)
			protected.GET("/servers", controllers.GetServersByUserID)

			protected.POST("/servers/:serverId/channels", controllers.CreateChannel)
			protected.GET("/servers/:serverId/channels", controllers.GetChannelsByServerID)
		}
	}

}
