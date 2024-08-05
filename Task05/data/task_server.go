package data

import (
	"context"
	"errors"
	"log"

	"github.com/legend123213/go_togo/Task05/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// AddTask adds a new task to the database.
func ServAddTask(task *models.Task, s *mongo.Database) (*models.Task, error) {
	store := s.Collection("Tasks")
	data, err := store.InsertOne(context.TODO(), *task)
	id := data.InsertedID.(primitive.ObjectID)
	task.ID = id
	return task, err
}

// GetTask retrieves a task from the database based on the given ID.
func ServGetTask(id string, s *mongo.Database) (*models.Task, error) {
	var task models.Task
	ID, _ := primitive.ObjectIDFromHex(id)
	err := s.Collection("Tasks").FindOne(context.TODO(), bson.M{"_id": ID}).Decode(&task)
	return &task, err
}

// GetTasks retrieves all tasks from the database.
func ServGetTasks(s *mongo.Database) ([]models.Task, error) {
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
func ServDeleteTask(id string, s *mongo.Database) error {
	ID, _ := primitive.ObjectIDFromHex(id)
	check, err := s.Collection("Tasks").DeleteOne(context.TODO(), bson.M{"_id": ID})
	if check.DeletedCount==0{
		return errors.New("")
	}
	return err
}

// EditTask updates a task in the database based on the given ID.
func ServEditTask(id string, s *mongo.Database, t *models.Task) (*models.Task, error) {
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

