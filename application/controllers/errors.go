package controllers

import "github.com/gin-gonic/gin"

func abortWithMessage(ctx *gin.Context, code int, message string) {
	ctx.AbortWithStatusJSON(code, map[string]interface{}{
		"message": message,
	})
}
