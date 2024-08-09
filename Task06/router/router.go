package router

import (
	"github.com/gin-gonic/gin"
	"github.com/legend123213/go_togo/Task06/controllers"
	"github.com/legend123213/go_togo/Task06/middleware"
	"go.mongodb.org/mongo-driver/mongo"
)


func Api(dbmongo *mongo.Database) *gin.Engine{
	server:=gin.New()
	server.Use(gin.Recovery())
	var userController controllers.UserControllerInter = controllers.NewUc(dbmongo)
	var taskController controllers.TaskControllerInter = controllers.NewTc(dbmongo)

	//user route
	server.POST("api/v1/login",userController.LogUser)
	server.POST("api/v1/signup",userController.CreateUser)
	server.Use(middleware.AuthMiddleware())
	server.PUT("api/v1/user/:id",userController.UpdateUser)
	server.GET("api/v1/task/:id",taskController.GetTask)
	server.GET("api/v1/tasks",taskController.GetAllTask)
	server.GET("api/v1/user/:id",userController.GetUser)
	server.Use(middleware.AdminMiddleware())
	
	server.GET("api/v1/users",userController.GetAllUser)
	
	server.DELETE("api/v1/user/:id",userController.RemoveUser)
	server.PATCH("api/v1/promote/:id",userController.MakeAdmin)


	server.POST("api/v1/task",taskController.CreateTask)
	server.PUT("api/v1/task/:id",taskController.UpdateTask)
	server.DELETE("api/v1/task/:id",taskController.RemoveTask)


	
	return server

}