package middleware

import (
	"net/http"

	"github.com/aamilineni/go-github/api/model"
	"github.com/gin-gonic/gin"
)

func ValidateJSONHeader(ctx *gin.Context) {

	acceptHeader := ctx.Request.Header.Get("Accept")
	if acceptHeader != "application/json" {
		ctx.JSON(http.StatusNotAcceptable, model.ErrorModel{
			Status:  http.StatusNotAcceptable,
			Message: "Header for Key `Accept` should have value of `application/json`",
		})
		ctx.Abort()
	}
}
