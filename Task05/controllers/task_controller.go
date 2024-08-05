package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/legend123213/go_togo/Task05/data"
	"github.com/legend123213/go_togo/Task05/models"
	"go.mongodb.org/mongo-driver/mongo"
)

// AddTasks is a controller function that adds a new task to the database
func AddTasks(c *gin.Context, storage *mongo.Database) {
	var task models.Task
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	data, err := data.ServAddTask(&task, storage)

	if err != nil {
		c.JSON(http.StatusInternalServerError, "db error")
		return
	}
	c.JSON(http.StatusAccepted, data)
}

// GetTask is a controller function that retrieves a specific task from the database
func GetTask(c *gin.Context, storage *mongo.Database) {
	id := c.Param("id")
	data, err := data.ServGetTask(id, storage)
	if err != nil {
		c.JSON(http.StatusBadRequest, "can't find the task")
		return
	}
	c.JSON(http.StatusOK, data)
}

// GetTasks is a controller function that retrieves all tasks from the database
func GetTasks(c *gin.Context, storage *mongo.Database) {
	data, err := data.ServGetTasks(storage)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}
	c.JSON(http.StatusOK, data)
}

// DeleteTask is a controller function that deletes a specific task from the database
func DeleteTask(c *gin.Context, storage *mongo.Database) {
	id := c.Param("id")
	err := data.ServDeleteTask(id, storage)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message":"no task found to be delete"})
		return
	}
	c.JSON(http.StatusNoContent, gin.H{"message": "Successfully Deleted"})
}

// EditTask is a controller function that edits a specific task in the database
func EditTask(c *gin.Context, storage *mongo.Database) {
	var task models.Task
	id := c.Param("id")
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	edited, errDb := data.ServEditTask(id, storage, &task)
	if errDb != nil {
		c.JSON(http.StatusNotFound,gin.H{"Message":"task not found to be edited"})
		return
	}
	c.JSON(http.StatusAccepted, edited)
}