package data

import (
	"context"
	"log"

	"github.com/legend123213/go_togo/Task05/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)










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
func DeleteTask(id string, s *mongo.Database) error {
	ID,_:=primitive.ObjectIDFromHex(id)
	_,err:=s.Collection("Tasks").DeleteOne(context.TODO(),bson.M{"_id":ID})
	return err
}
func EditTask(id string,s *mongo.Database,t *models.Task)(*models.Task,error){
	ID,_ := primitive.ObjectIDFromHex(id)
	update:=bson.M{
		"$set":bson.M{
			"title":t.Title,
			"description":t.Description,
			"due_date":t.Due_date,
			"status":t.Status,
		},
	}
	data,err := s.Collection("Tasks").UpdateOne(context.TODO(),bson.M{"_id":ID},update)
	t.ID = ID
	log.Println(err,data)
	return t,err

}