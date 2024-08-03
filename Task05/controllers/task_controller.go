package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/legend123213/go_togo/Task05/data"
	"github.com/legend123213/go_togo/Task05/models"
	"go.mongodb.org/mongo-driver/mongo"
)

func AddTasks(c *gin.Context,storage *mongo.Database){
	var task models.Task
	if err := c.ShouldBindJSON(&task); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
	data,err :=data.AddTask(&task, storage)

	if err!=nil{
		c.JSON(http.StatusInternalServerError,"db error")
		return
	}
	c.JSON(http.StatusAccepted,data)
}
func GetTask(c *gin.Context,storage *mongo.Database){
	// var task models.Task
	id := c.Param("id")
	data,err:=data.GetTask(id,storage)
	if err != nil {
		c.JSON(http.StatusBadRequest,"can't find the task")
		return 
	}
	c.JSON(http.StatusOK,data)

}