package data

import (
	"github.com/legend123213/go_togo/Task04/models"
)

type Storage struct{
	tasks map[int]models.Task
	ID_counter int
}

type TaskManager interface{
	AddTasks(task models.Task)
	GetTasks() []models.Task
	EditTask(id int,task models.Task)
	GetTask(id int) (models.Task,bool)
	DeleteTask(id int)
}

func DbRun() *Storage{
	return &Storage{
		tasks: make(map[int]models.Task),
		ID_counter: 1,
	}
}
func (ta *Storage) AddTasks(task models.Task){
	task.ID=ta.ID_counter
	ta.tasks[ta.ID_counter]= task
	ta.ID_counter++
}
func (ta *Storage) GetTasks() []models.Task{
	tasks := []models.Task{}
	for _,val := range ta.tasks{
		tasks = append(tasks,val)
	}
	return tasks
}
func (ta *Storage) GetTask(id int) (models.Task,bool){
	task,err := ta.tasks[id]
	return task,err
}
