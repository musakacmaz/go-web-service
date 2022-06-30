package routes

import (
	controllers "go-pet-api/src/controllers"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func DocumentationRoutes(router *gin.RouterGroup) {
	{
		router.GET("", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}
}

func HealthRoutes(router *gin.RouterGroup) {
	{
		router.GET("", controllers.GetHealth)
	}
}

func PetRoutes(router *gin.RouterGroup) {
	{
		router.GET("", controllers.GetPets)
		router.GET("/:id", controllers.GetPetsByID)
		router.POST("", controllers.CreatePet)
	}
}
