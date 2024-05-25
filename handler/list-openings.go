package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ListOpeninghandler(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"message": "GET openings",
	})
}
