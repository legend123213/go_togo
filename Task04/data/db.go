package data

import "github.com/legend123213/go_togo/Task04/models"

type Storage struct{
	tasks map[int]models.Task
	ID_counter int
}

type TaskManager interface{
	AddTasks(task models.Task)
	GetTasks(id ...int) []models.Task
	EditTask(id int,task models.Task)
	DeleteTask(id int)
}

func DbRun() *Storage{
	return &Storage{
		tasks: make(map[int]models.Task),
		ID_counter: 1,
	}
}
func (ta *Storage) AddTasks(task models.Task){
	ta.tasks[ta.ID_counter]= task
	
}
func (ta *Storage) GetTasks() []models.Task{
	tasks := []models.Task{}
	for _,val := range ta.tasks{
		tasks = append(tasks,val)
	}
	return tasks
}
