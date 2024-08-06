package router

import (
	"github.com/gin-gonic/gin"
	"github.com/legend123213/go_togo/Task06/controllers"
	"go.mongodb.org/mongo-driver/mongo"
)


func Api(dbmongo *mongo.Database) *gin.Engine{
	server:=gin.New()
	server.Use(gin.Recovery())
	var userController controllers.UserControllerInter = controllers.NewUc(dbmongo)
	var taskController controllers.TaskControllerInter = controllers.NewTc(dbmongo)

	//user route
	server.POST("api/v1/login",userController.CreateUser)
	server.GET("api/v1/user/:id",userController.GetUser)
	server.GET("api/v1/users",userController.GetAllUser)
	server.PUT("api/v1/user/:id",userController.UpdateUser)
	server.DELETE("api/v1/user/:id",userController.RemoveUser)


	//task route
	server.POST("api/v1/login",taskController.CreateTask)
	server.GET("api/v1/user/:id",taskController.GetTask)
	server.GET("api/v1/users",taskController.GetAllTask)
	server.PUT("api/v1/user/:id",taskController.UpdateTask)
	server.DELETE("api/v1/user/:id",taskController.RemoveTask)


	
	return server

}