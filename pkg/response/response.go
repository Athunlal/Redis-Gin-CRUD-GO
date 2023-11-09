package response

import "github.com/gin-gonic/gin"

func RespondWithError(ctx *gin.Context, statusCode int, err error) {
	ctx.JSON(statusCode, gin.H{
		"error": err.Error(),
	})
}

func RespondWithSuccess(ctx *gin.Context, message string, data any) {
	ctx.JSON(200, gin.H{
		"Message": message,
		"Data":    data,
	})
}
