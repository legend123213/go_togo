package data

import (
	"context"
	"log"

	"github.com/legend123213/go_togo/Task05/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)




type Tasks interface{
	AddTask(task models.Task) (models.Task,error)
	EditTask(task models.Task,id int) (models.Task,error)
	GetTask(id int) (models.Task,error)
	GetTasks() ([]models.Task,error)
	DeleteTask(id int) error
}





func AddTask(task *models.Task,s *mongo.Database) (*models.Task,error){
	store := s.Collection("Tasks")
	data,err:=store.InsertOne(context.TODO(),*task)
	id:=data.InsertedID.(primitive.ObjectID)
	task.ID=id
	return task,err
}

func GetTask(id string,s *mongo.Database) (*models.Task,error){
	var task models.Task
	ID,_ := primitive.ObjectIDFromHex(id)
	log.Println(ID)
	err := s.Collection("Tasks").FindOne(context.TODO(),bson.M{"_id": ID}).Decode(&task)
	log.Println(err)
	return &task,err
}
func GetTasks(s *mongo.Database)([]models.Task,error){
	var tasks []models.Task
	iterDocument,err := s.Collection("Tasks").Find(context.TODO(),bson.D{})
	for iterDocument.Next(context.TODO()){
		var task models.Task
		if err:= iterDocument.Decode(&task); err !=nil {
			return tasks,err
		}
		tasks = append(tasks, task)
	}
	log.Println(tasks)
	return tasks,err
}