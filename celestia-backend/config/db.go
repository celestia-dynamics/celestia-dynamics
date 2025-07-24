package config

import(
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DB *mongo.Client

func ConnectDB() *mongo.Client{
	mongoURI := "mongodb+srv://dynamicscelestia:jprxwIJ7wpja3h7z@celestiacluster.bbaluuc.mongodb.net/?retryWrites=true&w=majority&appName=CelestiaCluster"
	clientOptions := options.Client().ApplyURI(mongoURI)
	ctx,cancel := context.WithTimeout(context.Background(),10*time.Second)
	defer cancel();

	client,err := mongo.Connect(ctx,clientOptions)
	if err!=nil{
		log.Fatal(err);
	}

	err = client.Ping(ctx,nil)
	if err != nil{
		log.Fatal("Couldn't connect to MongoDB:",err)
	}

	fmt.Println("✅Connected to MongoDB Atlas")
	DB = client
	return client
}

func GetCollection(client *mongo.Client,collectionName string) *mongo.Collection{
	return client.Database("celestia").Collection(collectionName)
}