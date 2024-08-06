package data

import (
	"log"

	"github.com/legend123213/go_togo/Task06/models"
	"go.mongodb.org/mongo-driver/mongo"
)


type UserServices interface{
	AddUser(u *models.User,s *mongo.Database) error
	EditUser(user *models.User,s *mongo.Database) *models.User
	GetUser(id string,s *mongo.Database) *models.User
	DeleteUser(id string,s *mongo.Database) error
}

type UserServiceRepo struct {


}
func NewUserService() *UserServiceRepo{
	return &UserServiceRepo{

	}
}
func (u *UserServiceRepo)AddUser (user *models.User,s *mongo.Database) error{
	log.Println(u,s,user)
	return nil
}
func (u *UserServiceRepo)EditUser (user *models.User,s *mongo.Database) *models.User{
	log.Println(u,s,user)
	return nil
}
func (u *UserServiceRepo)GetUser (id string,s *mongo.Database) *models.User{
	var user models.User
	log.Println(id,s,user)
	return &models.User{}
}
func (u *UserServiceRepo)DeleteUser (id string,s *mongo.Database) error{
	
	log.Println(id,s,id)
	return nil
}