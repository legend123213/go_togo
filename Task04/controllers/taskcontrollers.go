package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/legend123213/go_togo/Task04/data"
	"github.com/legend123213/go_togo/Task04/models"
)
type inputTask struct{
	
	Title string `json:"Title" binding:"required"`
	Description string `json:"Description" binding:"required"`
	Due_date string `json:"Due_date" binding:"required"`
	Status string `json:"Status" binding:"required"`
}

func AddBook(c *gin.Context,storage data.TaskManager) {
	var task inputTask
	if err := c.BindJSON(&task); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data","erro":err})
			return
		}
	t := models.Task{
			Title: task.Title,
			Description: task.Description,
			Due_date: task.Due_date,
			Status: task.Status,
		}
	
	c.JSON(http.StatusCreated,storage.AddTasks(t))
	}



func GetTasks(c *gin.Context,storage data.TaskManager){
	task:= storage.GetTasks()
	fmt.Println(task)
	c.JSON(http.StatusOK, task)
}
func GetTask(c *gin.Context,storage data.TaskManager){
	Id:=c.Param("id")
	id,err:=strconv.Atoi(Id)
	task,exist:= storage.GetTask(id)
	if err!=nil{
		c.JSON(http.StatusNotFound,gin.H{"message":"wrong id"})
		
	}else if !exist{
		c.JSON(http.StatusBadRequest,gin.H{"message":"task not found"})
	}else{
c.JSON(http.StatusOK,task)
	}
}
func EditTask(c *gin.Context,storage data.TaskManager){
	Id:=c.Param("id")
	id,err:=strconv.Atoi(Id)
	var task inputTask
	if err := c.BindJSON(&task); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data","erro":err})
			return
		}
	t := models.Task{
			Title: task.Title,
			Description: task.Description,
			Due_date: task.Due_date,
			Status: task.Status,
		}
	editedTask,exist:= storage.EditTasks(id,t)
	if err!=nil{
		c.JSON(http.StatusNotFound,gin.H{"message":"wrong id"})
		
	}else if !exist{
		c.JSON(http.StatusBadRequest,gin.H{"message":"task not found"})
	}else{
	c.JSON(http.StatusOK,editedTask)
	}

}
func DeleteTask(c *gin.Context,storage data.TaskManager){
	Id:=c.Param("id")
	id,err:=strconv.Atoi(Id)
	exist:= storage.DeleteTask(id)
	if err!=nil{
		c.JSON(http.StatusNotFound,gin.H{"message":"wrong id"})
		
	}else if !exist{
		c.JSON(http.StatusBadRequest,gin.H{"message":"task not found"})
	}else{
c.JSON(http.StatusNoContent,gin.H{"message":"Successfully Deleted"})
	}

}


	

	

