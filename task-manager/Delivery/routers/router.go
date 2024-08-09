package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/legend123213/go_togo/Task07/task-manager/Delivery/controllers"
	Infrastructure "github.com/legend123213/go_togo/Task07/task-manager/Infrastructure"
	repositories "github.com/legend123213/go_togo/Task07/task-manager/Repositories"
	usecases "github.com/legend123213/go_togo/Task07/task-manager/UseCases"
	"go.mongodb.org/mongo-driver/mongo"
)


func Api(dbmongo *mongo.Database) *gin.Engine{
	server:=gin.New()
	tr := repositories.NewTaskService(dbmongo)
	ur:= repositories.NewUserService(dbmongo)
	var UserCase  usecases.UserUsecaseInt =usecases.NewuserUsecase(ur)
	var TaskCase usecases.TaskUseCaseint =  usecases.NewTaskUsecase(tr)
	var userController controllers.UserControllerInter = controllers.NewUc(UserCase)
	var taskController controllers.TaskControllerInter = controllers.NewTc(TaskCase)
	server.Use(gin.Recovery())
	server.Use(gin.ErrorLogger())
	server.POST("api/v1/login",userController.LogUser)
	server.POST("api/v1/signup",userController.CreateUser)
	server.Use(Infrastructure.AuthMiddleware())
	server.PUT("api/v1/user/:id",userController.UpdateUser)
	server.GET("api/v1/task/:id",taskController.GetTask)
	server.GET("api/v1/tasks",taskController.GetAllTask)
	server.GET("api/v1/user/:id",userController.GetUser)
	server.Use(Infrastructure.AdminMiddleware())
	
	server.GET("api/v1/users",userController.GetAllUser)
	
	server.DELETE("api/v1/user/:id",userController.RemoveUser)
	server.PATCH("api/v1/promote/:id",userController.MakeAdmin)


	server.POST("api/v1/task/",taskController.CreateTask)
	server.PUT("api/v1/task/:id",taskController.UpdateTask)
	server.DELETE("api/v1/task/:id",taskController.RemoveTask)


	
	return server

}