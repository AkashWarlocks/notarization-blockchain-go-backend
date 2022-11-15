package main

import (
	"example.com/main/db"
	"example.com/main/routes"
	"example.com/main/utils"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	db.ConnectDB()
	utils.ConnectHedera()
	routes.RouteIndex(router)
	router.Run("localhost:3000")
}