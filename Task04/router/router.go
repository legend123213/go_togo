package router

import (
	"github.com/gin-gonic/gin"
	"github.com/legend123213/go_togo/Task04/controllers"
	"github.com/legend123213/go_togo/Task04/data"
)

func Api() *gin.Engine{
	var storage data.TaskManager = data.DbRun()
	// gin.SetMode(gin.ReleaseMode)
	api:= gin.New()
	
	api.Use(gin.Recovery())
	api.GET("api/v1/task", func(c *gin.Context) {
		controllers.GetTasks(c,storage)
	})
	api.POST("api/v1/task",func(c *gin.Context){
		controllers.AddBook(c,storage)
	})
	api.GET("api/v1/task/:id",func(c *gin.Context){
		controllers.GetTask(c,storage)
	})
	api.PUT("api/v1/task/:id",func(c *gin.Context){
		controllers.EditTask(c,storage)
	})
	api.DELETE("api/v1/task/:id",func(c *gin.Context){
		controllers.DeleteTask(c,storage)
	})
	return api
}
