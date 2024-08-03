package router

import (
	"github.com/gin-gonic/gin"
	"github.com/legend123213/go_togo/Task05/controllers"
	"go.mongodb.org/mongo-driver/mongo"
)


func Api(storage *mongo.Database) *gin.Engine{
	api:=gin.New()
	api.Use(gin.Recovery())
	api.POST("api/v1/task", func(c *gin.Context){
		controllers.AddTasks(c,storage)
	} )
	api.GET("api/v1/task/:id",func(c *gin.Context){
		controllers.GetTask(c,storage)
	})
	
	
  return api

}