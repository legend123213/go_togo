package main

import (
	"context"
	"log"
	"os"

	"github.com/joho/godotenv"
	router "github.com/legend123213/go_togo/Task07/task-manager/Delivery/routers"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)



func init(){
	 err := godotenv.Load()
    if err != nil {
        log.Println("Error loading .env file")
        return
    }
}

func main(){
	MONGO_CLOUD_URL:=os.Getenv("MONGO_CLOUD_URL")
	dbmongo:=mongoService(MONGO_CLOUD_URL,"TaskManger")
	server:=router.Api(dbmongo)
	server.Run(":8000")

}


func mongoService(url string,dbname string) *mongo.Database{
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(url).SetServerAPIOptions(serverAPI)
	client,err := mongo.Connect(context.TODO(),opts)
	if err != nil {
		log.Println("mongo server error")
	}else{
		log.Println("Mongo connected")
	}
	DB:=client.Database(dbname)
	return DB
}