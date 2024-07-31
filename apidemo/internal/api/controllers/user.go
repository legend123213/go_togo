package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Getuser(c *gin.Context) {
	 c.IndentedJSON(http.StatusOK, gin.H{"message": "hello there"})
}
func Adduser(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{"message": "added succefully"})
}
