package main

import (
	_ "go-pet-api/docs"
	"go-pet-api/src/routes"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @title Pet API
// @version 1.0
// @description This is a sample web service.

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name MIT
// @license.url https://opensource.org/licenses/MIT

// @host localhost:8080
// @BasePath /
// @query.collection.format multi

func main() {
	app := gin.Default()
	app.SetTrustedProxies(nil)
	app.NoRoute(func(ctx *gin.Context) {
		ctx.AbortWithStatus(http.StatusNotFound)
	})
	petRoutes := app.Group("/pets")
	healthRoutes := app.Group("/health")
	documentationRoutes := app.Group("/documentation/*any")
	routes.PetRoutes(petRoutes)
	routes.HealthRoutes(healthRoutes)
	routes.DocumentationRoutes(documentationRoutes)

	app.Run() // listen and serve on default port
}
