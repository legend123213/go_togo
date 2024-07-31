package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/legend123213/go_togo/tree/master/apidemo/internal/api/controllers"
)

func Setup() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	api := gin.New()
	api.Use(gin.Recovery())
	api.GET("/api/v1/user", controllers.Getuser)
	api.POST("api/v1/user", controllers.Adduser)
	return api

}
