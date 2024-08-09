package controllers

import (
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
func (u *Uc)CreateUser(c *gin.Context) {
	var user domain.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := u.userusecase.IsUsernameUnique(user.Username); err != nil {
		c.JSON(http.StatusConflict, gin.H{"message": err.Error()})
		return 
	}

	res, err := u.userusecase.Register(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": res})
		return 
	}
	c.JSON(http.StatusCreated, gin.H{"message": "User registered successfully", "token": res})
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
	c.JSON(http.StatusOK, gin.H{"message": "User login in successfully", "token": res})

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
	c.JSON(http.StatusOK, gin.H{"message": "User profile updated successfully", "user": editedUser})

}
func (u *Uc)GetUser(c *gin.Context){
	id := c.Param("id")
	user,err := u.userusecase.Fetch(id)
	if err != nil{
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "User profile fetched successfully", "user": user})
}

func (u *Uc)RemoveUser(c *gin.Context){
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user_id must be added in the request"})
		return
	}
	
	err := u.userusecase.Delete(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}
func (u *Uc)GetAllUser(c *gin.Context){
	service := u.userusecase.FetchAllUser()
	c.JSON(http.StatusOK, gin.H{"message": "All users fetched successfully", "users": service})
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
	c.JSON(http.StatusOK, gin.H{"message": "User promoted to admin successfully"})
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
		c.JSON(http.StatusBadRequest, gin.H{"error": "user_id must be added in the request"})
		return
	}
	data, err := u.taskusecase.Create(&task)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create task"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Task created successfully", "task": data})

}
func (u *Tc)UpdateTask(c *gin.Context){
	var task domain.Task
	id := c.Param("id")
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	edited, errDb := u.taskusecase.UpdateTask(id, &task)
	if errDb != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Task not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Task updated successfully", "task": edited})
}
func (u *Tc)GetTask(c *gin.Context){
	id := c.Param("id")
	admin := c.MustGet("isActive").(bool)
	user_id:=c.MustGet("id").(primitive.ObjectID)
	if admin{
		data, err := u.taskusecase.FetchTask(id, "",)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error":"Task not found"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "Task fetched successfully", "task": data})
	}else{
		user_ID:=user_id.String()
		data, err := u.taskusecase.FetchTask(id,user_ID)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error":"Task not found"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "Task fetched successfully", "task": data})
	}
}
func (u *Tc)RemoveTask(c *gin.Context){
	id := c.Param("id")
	err := u.taskusecase.RemoveTask(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "No task found to be deleted"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Task deleted successfully"})
}
func (u *Tc)GetAllTask(c *gin.Context){
	admin := c.MustGet("isActive").(bool)
	user_id:=c.MustGet("id").(primitive.ObjectID)
	if admin{
		data, err := u.taskusecase.FetchTasks("")
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "All tasks fetched successfully", "tasks": data})
	}else{
		user_ID:=user_id.String()
		data, err := u.taskusecase.FetchTasks(user_ID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "User tasks fetched successfully", "tasks": data})
	}
}
