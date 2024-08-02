package data

import (
	"github.com/legend123213/go_togo/Task04/models"
)

type Storage struct{
	tasks map[int]models.Task
	ID_counter int
}

type TaskManager interface{
	AddTasks(task models.Task) models.Task
	GetTasks() []models.Task
	EditTasks(id int,task models.Task) (models.Task,bool)
	GetTask(id int) (models.Task,bool)
	DeleteTask(id int) bool
}

func DbRun() *Storage{
	return &Storage{
		tasks: make(map[int]models.Task),
		ID_counter: 1,
	}
}
func (ta *Storage) AddTasks(task models.Task) models.Task{
	task.ID=ta.ID_counter
	ta.tasks[ta.ID_counter]= task
	ta.ID_counter++
	return task
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

func (ta *Storage) EditTasks(id int,t models.Task) (models.Task,bool) {
	t.ID = id
	_,exist:=ta.tasks[id]
	if exist{
		ta.tasks[id]=t
	}
return ta.tasks[id],exist
}
func (ta *Storage) DeleteTask(id int) bool{
	_,exist:=ta.tasks[id]
	if exist{
		delete(ta.tasks,id)
	}
	return exist
}