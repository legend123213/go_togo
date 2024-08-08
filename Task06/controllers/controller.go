package controllers

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/legend123213/go_togo/Task06/data"
	"github.com/legend123213/go_togo/Task06/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)
func isUsernameUnique(ctx context.Context, collection *mongo.Collection, username string) error {
    var existingUser models.User
    err := collection.FindOne(ctx, bson.M{"username": username}).Decode(&existingUser)
    if err == nil {
        return fmt.Errorf("username '%s' already exists", username)
    }
    if err == mongo.ErrNoDocuments {
        return nil 
		}

    return err
}
type UserControllerInter interface{
	CreateUser(c *gin.Context)
	GetUser(c *gin.Context)
	GetAllUser(c *gin.Context)
	UpdateUser(c *gin.Context)
	RemoveUser(c *gin.Context)
	LogUser(c *gin.Context)
	MakeAdmin(c *gin.Context)
}

type Uc struct {
	 DbStorage  *mongo.Database
	 service data.UserServices
}

func NewUc(db *mongo.Database) *Uc{
		var u data.UserServices = data.NewUserService()
	return &Uc{
		DbStorage:db,
		service :u,
	}
}
func (u *Uc)CreateUser(c *gin.Context){
	storage := u.DbStorage
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := isUsernameUnique(context.TODO(), u.DbStorage.Collection("Users"), user.Username); err != nil {
        c.JSON(http.StatusConflict,gin.H{"message":err.Error() })
		  return 
    }

	res,err := u.service.RegisterUser(&user,storage)
	if err!=nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": res})
	}
	c.JSON(http.StatusAccepted, gin.H{"message": "User registered in successfully", "token": res})

}
func (u *Uc) LogUser(c *gin.Context){
	storage := u.DbStorage
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res,err := u.service.LoginUser(&user,storage)
	if err!=nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": res})
		return 
	}
	c.JSON(http.StatusAccepted, gin.H{"message": "User login in successfully", "token": res})

}
func (u *Uc)UpdateUser(c *gin.Context){

}
func (u *Uc)GetUser(c *gin.Context){
	id := c.Param("id")
	store := u.DbStorage
	user,err := u.service.GetUser(id,store)
	if err != nil{
		c.JSON(http.StatusNotFound,err)
		return
	}
	c.JSON(http.StatusOK,user)

}
func (u *Uc)RemoveUser(c *gin.Context){

}
func (u *Uc)GetAllUser(c *gin.Context){
	store:=u.DbStorage
	service := u.service.GetAllUser(store)
	c.JSON(http.StatusOK,service)

}
func (u *Uc)MakeAdmin(c *gin.Context){
	id := c.Param("id")
	store:=u.DbStorage
	err := u.service.RoleChanger(id,store)
	if err != nil {
		if err.Error() == "user already admin" {
			c.JSON(http.StatusConflict, gin.H{"message": "User is already an admin"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}
	c.JSON(http.StatusAccepted,gin.H{"message":"user promoted"})

}



type TaskControllerInter interface{
	CreateTask(c *gin.Context)
	GetTask(c *gin.Context)
	GetAllTask(c *gin.Context)
	UpdateTask(c *gin.Context)
	RemoveTask(c *gin.Context)
}

type Tc struct {
	 DbStorage  *mongo.Database
	 task data.TaskInterface
}

func NewTc(db *mongo.Database) *Tc{
	var t data.TaskInterface = data.NewTaskService()
	return &Tc{
		DbStorage:db,
		task:t,
	}
}
func (u *Tc)CreateTask(c *gin.Context){
	storage :=u.DbStorage
	var task models.Task
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	dataaa, _ := c.Get("id")
	task.UserID = dataaa.(primitive.ObjectID)
	data, err_ := u.task.SAddTask(&task, storage)

	if err_ != nil {
		c.JSON(http.StatusInternalServerError, "db error")
		return
	}
	c.JSON(http.StatusAccepted, data)

}
func (u *Tc)UpdateTask(c *gin.Context){
	storage := u.DbStorage
	var task models.Task
	id := c.Param("id")
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	edited, errDb := u.task.SEditTask(id, storage, &task)
	if errDb != nil {
		c.JSON(http.StatusNotFound,gin.H{"Message":"task not found to be edited"})
		return
	}
	c.JSON(http.StatusAccepted, edited)

}
func (u *Tc)GetTask(c *gin.Context){
	storage := u.DbStorage
	id := c.Param("id")
	admin := c.MustGet("isActive").(bool)
	user_id:=c.MustGet("id").(primitive.ObjectID)
	if admin{
		data, err := u.task.SGetTask(id, "",storage)
		if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message":"can't find the task"})
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, data)
	}else{
		user_ID:=user_id.String()
		data, err := u.task.SGetTask(id,user_ID , storage)
		if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message":"can't find the task"})
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, data)
	}

	


}
func (u *Tc)RemoveTask(c *gin.Context){
	storage:= u.DbStorage
	id := c.Param("id")
	err :=u.task.SDeleteTask(id, storage)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message":"no task found to be delete"})
		return
	}
	c.JSON(http.StatusNoContent, gin.H{"message": "Successfully Deleted"})

}
func (u *Tc)GetAllTask(c *gin.Context){
	storage := u.DbStorage
	admin := c.MustGet("isActive").(bool)
	user_id:=c.MustGet("id").(primitive.ObjectID)
	if admin{
		data, err := u.task.SGetTasks("",storage)
		log.Println(data,err)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, data)
		
	}else{
		user_ID:=user_id.String()
		data, err := u.task.SGetTasks(user_ID,storage)
		if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, data)
	}
	

}