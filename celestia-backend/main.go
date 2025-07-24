package main

import (
	"celestia-backend/config"
	"celestia-backend/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	config.ConnectDB() // 🔌 Connects to DB

	r := gin.Default()
	

	// 🧠 Pass the DB from config
	routes.RegisterRoutes(r, config.DB)

	r.Run(":8080")
}
