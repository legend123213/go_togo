package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/legend123213/go_togo/Task05/router"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)
var MONGO_CLOUD_URL string


func init(){
	 err := godotenv.Load()
    if err != nil {
        fmt.Println("Error loading .env file")
        return
    }
}
func main(){
	//importing api key for cloud mongo1
	MONGO_CLOUD_URL=os.Getenv("MONGO_CLOUD_URL")
	
	//initializer for mongo server
	client := MongoStarter("TaskManger")

	// initializer for mongo server
	server:=router.Api(client)
	server.Run(":8000")

}

//function that connect mongodb and return  database instance
func MongoStarter(dbname string) *mongo.Database{
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(MONGO_CLOUD_URL).SetServerAPIOptions(serverAPI)
	client,err := mongo.Connect(context.TODO(),opts)
	if err != nil {
		log.Println("mongo server error")
	}else{
		log.Println("Mongo connected")
	}
	DB:=client.Database(dbname)
	return DB
}