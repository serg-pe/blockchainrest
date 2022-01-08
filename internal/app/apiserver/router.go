package apiserver

import (
	"github.com/gin-gonic/gin"
	"github.com/serg-pe/blockchainrest/internal/app/apiserver/controllers"
)

func NewRouter() *gin.Engine {
	router := gin.New()
	helloController := controllers.HelloController{}
	router.GET("/hello", helloController.HelloWorld)
	return router
}
