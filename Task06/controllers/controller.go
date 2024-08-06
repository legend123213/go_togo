package controllers

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserControllerInter interface{
	CreateUser(c *gin.Context)
	GetUser(c *gin.Context)
	GetAllUser(c *gin.Context)
	UpdateUser(c *gin.Context)
	RemoveUser(c *gin.Context)
}

type Uc struct {
	 Db  *mongo.Database
}

func NewUc(db *mongo.Database) *Uc{
	return &Uc{
		Db:db,
	}
}
func (u *Uc)CreateUser(c *gin.Context){

}
func (u *Uc)UpdateUser(c *gin.Context){

}
func (u *Uc)GetUser(c *gin.Context){

}
func (u *Uc)RemoveUser(c *gin.Context){

}
func (u *Uc)GetAllUser(c *gin.Context){

}



type TaskControllerInter interface{
	CreateTask(c *gin.Context)
	GetTask(c *gin.Context)
	GetAllTask(c *gin.Context)
	UpdateTask(c *gin.Context)
	RemoveTask(c *gin.Context)
}

type Tc struct {
	 Db  *mongo.Database
}

func NewTc(db *mongo.Database) *Uc{
	return &Uc{
		Db:db,
	}
}
func (u *Uc)CreateTask(c *gin.Context){

}
func (u *Uc)UpdateTask(c *gin.Context){

}
func (u *Uc)GetTask(c *gin.Context){

}
func (u *Uc)RemoveTask(c *gin.Context){

}
func (u *Uc)GetAllTask(c *gin.Context){

}