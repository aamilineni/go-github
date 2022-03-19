package router

import (
	"github.com/gin-gonic/gin"

	"github.com/aamilineni/go-github/api/handlers/github"
)

func InitialiseRouter() *gin.Engine {

	// Creates a gin router with default middleware:
	// logger and recovery (crash-free) middleware
	router := gin.Default()

	router.GET("/:name/repos", github.NewGithubHandler(router).Get)

	return router

}
