package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type HelloController struct{}

func (controller HelloController) HelloWorld(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"message": "Hello, Go!",
	})
}
