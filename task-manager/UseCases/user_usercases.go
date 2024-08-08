package usecases

import (
	domain "github.com/legend123213/go_togo/Task07/task-manager/Domain"
	repositories "github.com/legend123213/go_togo/Task07/task-manager/Repositories"
)


type UserUsecaseInt interface{
	Register(u *domain.User) (string,error)
	Edit(id string,user *domain.User)(*domain.User,error)
	Fetch(id string) (*domain.User,error)
	Delete(id string) error
	Login(u *domain.User) (string,error)
	FetchUserByUname(username string)(*domain.User,error)
	RoleChanger(id string)(error)
	FetchAllUser() *[]domain.User
	IsUsernameUnique(username string) error


}

type UserUsecase struct {
	userRepository repositories.UserServices
}

func NewuserUsecase(userRepository repositories.UserServices) *UserUsecase {
	return &UserUsecase{
		userRepository: userRepository,
	}
}

func (use *UserUsecase)Register(u *domain.User) (string,error){
	return use.userRepository.RegisterUser(u)

}
func (use *UserUsecase)Edit(id string,user *domain.User)(*domain.User,error){
	return use.userRepository.EditUser(id,user)

}
func (use *UserUsecase)Fetch(id string) (*domain.User,error){
	return use.userRepository.GetUser(id)

}
func (use *UserUsecase)Delete(id string) error{
	return use.userRepository.DeleteUser(id)

}
func (use *UserUsecase)Login(u *domain.User) (string,error){
	return use.userRepository.LoginUser(u)

}
func (use *UserUsecase)FetchUserByUname(username string)(*domain.User,error){
	return use.userRepository.GetUserByUname(username)

}
func (use *UserUsecase)	RoleChanger(id string)(error){
	return use.userRepository.RoleChanger(id)

}
func (use *UserUsecase)	FetchAllUser() *[]domain.User{
	return use.userRepository.GetAllUser()

}
func (use *UserUsecase)	IsUsernameUnique(username string) error{
	return use.userRepository.IsUsernameUnique(username)

}

	
	


