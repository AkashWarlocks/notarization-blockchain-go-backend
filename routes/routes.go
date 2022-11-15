package routes

import (
	"example.com/main/controller"

	"github.com/gin-gonic/gin"
)

func RouteIndex(Router *gin.Engine)  {
	Router.POST("/setData", controller.SetData)
}