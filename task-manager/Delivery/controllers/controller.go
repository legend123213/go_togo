package controllers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	domain "github.com/legend123213/go_togo/Task07/task-manager/Domain"
	usecases "github.com/legend123213/go_togo/Task07/task-manager/UseCases"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

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
	 userusecase usecases.UserUsecaseInt
}

func NewUc(userusecase usecases.UserUsecaseInt) *Uc{
	return &Uc{
		userusecase:userusecase,
	}
}
func (u *Uc)CreateUser(c *gin.Context){
	var user domain.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := u.userusecase.IsUsernameUnique(user.Username); err != nil {
        c.JSON(http.StatusConflict,gin.H{"message":err.Error() })
		  return 
    }

	res,err := u.userusecase.Register(&user)
	if err!=nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": res})
	}
	c.JSON(http.StatusAccepted, gin.H{"message": "User registered in successfully", "token": res})

}
func (u *Uc) LogUser(c *gin.Context){
	var user domain.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res,err := u.userusecase.Login(&user)
	if err!=nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": res})
		return 
	}
	c.JSON(http.StatusAccepted, gin.H{"message": "User login in successfully", "token": res})

}
func (u *Uc)UpdateUser(c *gin.Context){
	var user domain.User
	ID:=c.Param("id")
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	editedUser,err := u.userusecase.Edit(ID,&user)
	if err!=nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK,editedUser)

}
func (u *Uc)GetUser(c *gin.Context){
	id := c.Param("id")
	user,err := u.userusecase.Fetch(id)
	if err != nil{
		c.JSON(http.StatusNotFound,err)
		return
	}
	c.JSON(http.StatusOK,user)

}

func (u *Uc)RemoveUser(c *gin.Context){
	id := c.Param("id")
	if id ==""{
		c.JSON(http.StatusBadRequest, gin.H{"error": "user_id must be add in request"})
		return
	}
	
	err:=u.userusecase.Delete(id)
	if err != nil{
		c.JSON(http.StatusNotFound,err)
		return
	}
	c.JSON(http.StatusAccepted,gin.H{"message":"deletion successful"})

}
func (u *Uc)GetAllUser(c *gin.Context){
	service := u.userusecase.FetchAllUser()
	c.JSON(http.StatusOK,service)

}
func (u *Uc)MakeAdmin(c *gin.Context){
	id := c.Param("id")
	err := u.userusecase.RoleChanger(id)
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
	 taskusecase usecases.TaskUseCaseint
}

func NewTc(taskUseCase usecases.TaskUseCaseint) *Tc{
	return &Tc{
		taskusecase:taskUseCase,
	}
}
func (u *Tc)CreateTask(c *gin.Context){
	var task domain.Task
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	objectID,_ := primitive.ObjectIDFromHex("")
	if task.UserID == objectID{
		c.JSON(http.StatusBadRequest, gin.H{"error": "user_id must be add in request"})
		return
	}
	data, err_ := u.taskusecase.Create(&task)

	if err_ != nil {
		c.JSON(http.StatusInternalServerError, "db error")
		return
	}
	c.JSON(http.StatusAccepted, data)

}
func (u *Tc)UpdateTask(c *gin.Context){
	var task domain.Task
	id := c.Param("id")
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	edited, errDb := u.taskusecase.UpdateTask(id, &task)
	if errDb != nil {
		c.JSON(http.StatusNotFound,gin.H{"Message":"task not found to be edited"})
		return
	}
	c.JSON(http.StatusAccepted, edited)

}
func (u *Tc)GetTask(c *gin.Context){
	id := c.Param("id")
	admin := c.MustGet("isActive").(bool)
	user_id:=c.MustGet("id").(primitive.ObjectID)
	if admin{
		data, err := u.taskusecase.FetchTask(id, "",)
		if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message":"can't find the task"})
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, data)
	}else{
		user_ID:=user_id.String()
		data, err := u.taskusecase.FetchTask(id,user_ID)
		if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message":"can't find the task"})
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, data)
	}

	


}
func (u *Tc)RemoveTask(c *gin.Context){
	id := c.Param("id")
	err :=u.taskusecase.RemoveTask(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message":"no task found to be delete"})
		return
	}
	c.JSON(http.StatusNoContent, gin.H{"message": "Successfully Deleted"})

}
func (u *Tc)GetAllTask(c *gin.Context){
	admin := c.MustGet("isActive").(bool)
	user_id:=c.MustGet("id").(primitive.ObjectID)
	if admin{
		data, err := u.taskusecase.FetchTasks("",)
		log.Println(data,err)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, data)
		
	}else{
		user_ID:=user_id.String()
		data, err := u.taskusecase.FetchTasks(user_ID)
		if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, data)
	}
	

}