package usecases

import (
	domain "github.com/legend123213/go_togo/Task08/task-manager/Domain"
	repositories "github.com/legend123213/go_togo/Task08/task-manager/Repositories"
)

type taskUsecase struct {
	taskRepository repositories.TaskInterface
}
//go:generate mockery --name TaskInterface    
type TaskUseCaseint interface {
	Create(task *domain.Task) (*domain.Task, error)
	FetchTasks(user_id string) ([]domain.Task, error)
	FetchTask(id string, username string) (*domain.Task, error)
	RemoveTask(id string) error
	UpdateTask(id string, task *domain.Task) (*domain.Task, error)
}

func NewTaskUsecase(taskRepository repositories.TaskInterface) *taskUsecase {
	return &taskUsecase{
		taskRepository: taskRepository,
	}
}
func (u *taskUsecase) Create(task *domain.Task) (*domain.Task, error) {
	return u.taskRepository.SAddTask(task)

}
func (u *taskUsecase) FetchTasks(user_id string) ([]domain.Task, error) {
	return u.taskRepository.SGetTasks(user_id)

}
func (u *taskUsecase) FetchTask(id string, username string) (*domain.Task, error) {
	return u.taskRepository.SGetTask(id,username)

}
func (u *taskUsecase) RemoveTask(id string) error {
	return u.taskRepository.SDeleteTask(id)

}
func (u *taskUsecase) UpdateTask(id string, task *domain.Task) (*domain.Task, error) {
	return u.taskRepository.SEditTask(id,task)

}
