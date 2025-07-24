package main

import (
	"celestia-backend/config"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main(){

	config.ConnectDB()
	router := gin.Default()

	router.GET("/",func(c *gin.Context){
		c.JSON(http.StatusOK,gin.H{
			"message": "Celestia Dynamics Backend is Live🚀",
		})
	})

	//TODO: Plug in routes here
	//Example: authRoutes(router)

	log.Println("Starting server on: 8080...")
	err := router.Run(":8080")
	if err!= nil{
		log.Fatal("Server couldn't start:", err)
	}
}