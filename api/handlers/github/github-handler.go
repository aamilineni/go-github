package github

import (
	"github.com/gin-gonic/gin"
)

type githubHandler struct {
}

func NewGithubHandler(c *gin.Engine) *githubHandler {
	return &githubHandler{}
}

func (me *githubHandler) Get(ctx *gin.Context) {
	name := ctx.Param("name")
	ctx.Bind(name)
}
