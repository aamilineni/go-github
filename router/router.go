package router

import (
	"github.com/gin-gonic/gin"

	"github.com/aamilineni/go-github/api/handlers"
	"github.com/aamilineni/go-github/restclient"
)

func InitialiseRouter() *gin.Engine {

	gin.SetMode(gin.ReleaseMode)

	// Creates a gin router with default middleware:
	// logger and recovery (crash-free) middleware
	router := gin.Default()

	router.GET("/:name/repos", handlers.NewGithubHandler(restclient.Client).Get)

	return router

}
