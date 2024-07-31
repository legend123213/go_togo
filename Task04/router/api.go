package router

import (
	"github.com/gin-gonic/gin"
	"github.com/legend123213/go_togo/Task04/controllers"
	"github.com/legend123213/go_togo/Task04/data"
)

func Api() *gin.Engine{
	var storage = data.DbRun()
	// gin.SetMode(gin.ReleaseMode)
	api:= gin.New()
	
	api.Use(gin.Recovery())
	api.GET("api/task", func(c *gin.Context) {
		controllers.GetTasks(c,storage)
	})
	api.POST("api/task",func(c *gin.Context){
		controllers.AddBook(c,storage)
	})
	api.GET("api/task/:id",func(c *gin.Context){
		controllers.GetTask(c,storage)
	})
	return api
}