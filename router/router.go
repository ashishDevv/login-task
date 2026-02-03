package router

import (
	"github.com/aashisDevv/login-api/controller"
	"github.com/gin-gonic/gin"
)

func SetupRouter(handler *controller.UserController) *gin.Engine {
	r := gin.Default()
	
	r.POST("/login", handler.Login)
	
	return r
}