package routes

import(
	"kordchat-be/controllers"
	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetupRoutes(r *gin.Engine){

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))


	api := r.Group("api")
	{
		api.POST("/register", controllers.Register)
		api.POST("/login", controllers.Login)
		api.POST("/refresh", controllers.RefreshToken)
	}
}