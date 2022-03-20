package router

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/aamilineni/go-github/api/handlers"
	"github.com/aamilineni/go-github/api/middleware"
	_ "github.com/aamilineni/go-github/docs"
	"github.com/aamilineni/go-github/restclient"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitialiseRouter() *gin.Engine {

	gin.SetMode(gin.ReleaseMode)

	// Creates a gin router with default middleware:
	// logger and recovery (crash-free) middleware
	router := gin.Default()

	// Routes
	// version 1
	apiV1 := router.Group("/api/v1")

	url := ginSwagger.URL("http://localhost:8080/swagger/doc.json") // The url pointing to API definition
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	// health check API
	apiV1.GET("/healthcheck", healthcheck)

	// get repos information API
	apiV1.GET("/:name/repos", middleware.ValidateJSONHeader, handlers.NewGithubHandler(restclient.Client).Get)

	return router

}

// HealthCheck godoc
// @Summary Show the status of server.
// @Description get the status of server.
// @Tags root
// @Accept */*
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router / [get]
func healthcheck(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, map[string]interface{}{
		"data": "Server is up and running",
	})
}
