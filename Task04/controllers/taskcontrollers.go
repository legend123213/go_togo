package controllers

import (
	"fmt"
	"log"
	"net/http"

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

func AddBook(c *gin.Context,storage *data.Storage) {
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
	storage.AddTasks(t)
	log.Println(t,storage)
	c.JSON(http.StatusCreated,gin.H{"message":"Succuss"})
	}



func GetTasks(c *gin.Context,storage *data.Storage){
	task:= storage.GetTasks()
	fmt.Println(task)
	log.Println("hello",task)
	c.JSON(http.StatusOK, task)
}