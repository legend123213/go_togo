package data

import (
	"context"
	"errors"
	"log"

	"github.com/legend123213/go_togo/Task06/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)



type TaskInterface interface{
	SAddTask(task *models.Task,s *mongo.Database)(*models.Task,error)
	SGetTask(id string,s *mongo.Database)(*models.Task, error)
	SGetTasks(s *mongo.Database) ([]models.Task, error)
	SDeleteTask(id string,s *mongo.Database) error
	SEditTask(id string, s *mongo.Database, t *models.Task) (*models.Task, error)

}

type TaskServiceRepo struct{

}

func NewTaskService() *TaskServiceRepo{
	return &TaskServiceRepo{
		
	}
}

// AddTask adds a new task to the database.
func (this *TaskServiceRepo)SAddTask(task *models.Task, s *mongo.Database) (*models.Task, error) {
	store := s.Collection("Tasks")
	data, err := store.InsertOne(context.TODO(), *task)
	id := data.InsertedID.(primitive.ObjectID)
	task.ID = id
	return task, err
}
// GetTask retrieves a task from the database based on the given ID.
func  (this *TaskServiceRepo)SGetTask(id string, s *mongo.Database) (*models.Task, error) {
	var task models.Task
	ID, _ := primitive.ObjectIDFromHex(id)
	err := s.Collection("Tasks").FindOne(context.TODO(), bson.M{"_id": ID}).Decode(&task)
	return &task, err
}

// GetTasks retrieves all tasks from the database.
func (this *TaskServiceRepo) SGetTasks(s *mongo.Database) ([]models.Task, error) {
	var tasks []models.Task
	iterDocument, err := s.Collection("Tasks").Find(context.TODO(), bson.D{})
	for iterDocument.Next(context.TODO()) {
		var task models.Task
		if err := iterDocument.Decode(&task); err != nil {
			return tasks, err
		}
		tasks = append(tasks, task)
	}
	return tasks, err
}

// DeleteTask deletes a task from the database based on the given ID.
func (this *TaskServiceRepo) SDeleteTask(id string, s *mongo.Database) error {
	ID, _ := primitive.ObjectIDFromHex(id)
	check, err := s.Collection("Tasks").DeleteOne(context.TODO(), bson.M{"_id": ID})
	if check.DeletedCount==0{
		return errors.New("")
	}
	return err
}

// EditTask updates a task in the database based on the given ID.
func  (this *TaskServiceRepo)SEditTask(id string, s *mongo.Database, t *models.Task) (*models.Task, error) {
	ID, err_:= primitive.ObjectIDFromHex(id)
	log.Println(err_)
	update := bson.M{
		"$set": bson.M{
			"title":       t.Title,
			"description": t.Description,
			"due_date":    t.Due_date,
			"status":      t.Status,
		},
	}
	check, err := s.Collection("Tasks").UpdateOne(context.TODO(), bson.M{"_id": ID}, update)
	t.ID = ID
	if check.MatchedCount==0{
		return t,errors.New("")
	}
	return t, err
}