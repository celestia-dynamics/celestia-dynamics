package routes

import (
	"celestia-backend/controllers"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func RegisterRoutes(r *gin.Engine, db *mongo.Database) {
	// Inject DB into the controller
	ac := controllers.NewAuthController(db)

	auth := r.Group("/api/auth")
	{
		auth.POST("/signup", ac.Signup)
		auth.POST("/login", ac.Login)
	}
}
