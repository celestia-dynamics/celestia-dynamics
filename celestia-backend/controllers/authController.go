package controllers

import (
	"context"
	"net/http"
	"strings"
	"time"

	"celestia-backend/config"
	"celestia-backend/models"
	"celestia-backend/utils"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

var userCollection = config.GetCollection(config.DB,"users")

func Signup(c *gin.Context){
	var user models.User

	if err:= c.ShouldBindJSON(&user); err!=nil{
		c.JSON(http.StatusBadRequest,gin.H{"error":err.Error()})
		return
	}

	ctx,cancel := context.WithTimeout(context.Background(),10*time.Second)
	defer cancel()

	count,_ := userCollection.CountDocuments(ctx,bson.M{"email":strings.ToLower(user.Email)})
	if count>0{
		c.JSON(http.StatusConflict,gin.H{"error":"Email already exists"})
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password),14)
	if err!=nil{
		c.JSON(http.StatusInternalServerError,gin.H{"error":"Failed to hash password"})
		return
	}

	user.Password = string(hashedPassword)
	user.Email = strings.ToLower((user.Email))
	user.CreatedAt = time.Now()
	user.ID = primitive.NewObjectID()

	_,err = userCollection.InsertOne(ctx,user)
	if err!=nil{
		c.JSON(http.StatusInternalServerError,gin.H{"error":"Failed to create user"})
		return
	}

	c.JSON(http.StatusCreated,gin.H{"message":"Signup successful"})
}

func Login(c *gin.Context){
	var input models.User
	var dbUser models.User

	if err := c.ShouldBindJSON(&input); err!=nil{
		c.JSON(http.StatusBadRequest,gin.H{"error":err.Error()})
		return
	}

	ctx,cancel := context.WithTimeout(context.Background(),10*time.Second)
	defer cancel()

	err:= userCollection.FindOne(ctx,bson.M{"email":strings.ToLower(input.Email)}).Decode(&dbUser)
	if err != nil{
		c.JSON(http.StatusUnauthorized,gin.H{"error":"Invalid email or password"})
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(dbUser.Password),[]byte(input.Password))
	if err != nil{
		c.JSON(http.StatusUnauthorized,gin.H{"error":"Invalid email or password"})
		return
	}

	token,err := utils.GenerateToken(dbUser.ID.Hex(),dbUser.Email)
	if err!= nil{
		c.JSON(http.StatusInternalServerError,gin.H{"error":"Failed to generate token"})
		return 
	}

	c.JSON(http.StatusOK,gin.H{
		"message":"Login successful",
		"token": token,
		"name": dbUser.Name,
		"email":dbUser.Email,
	})
}
