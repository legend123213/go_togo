package repositories

import (
	"context"
	"errors"

	domain "github.com/legend123213/go_togo/Task08/task-manager/Domain"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)



type TaskInterface interface{
	SAddTask(task *domain.Task)(*domain.Task,error)
	SGetTask(id string,username string)(*domain.Task, error)
	SGetTasks(user_id string) ([]domain.Task, error)
	SDeleteTask(id string) error
	SEditTask(id string, t *domain.Task) (*domain.Task, error)

}
type Collectioninter interface{
	InsertOne(ctx context.Context, document interface{}, opts ...*options.InsertOneOptions) (*mongo.InsertOneResult, error) 
	FindOne(ctx context.Context, filter interface{}, opts ...*options.FindOneOptions) *mongo.SingleResult
	Find(ctx context.Context, filter interface{}, opts ...*options.FindOptions) (*mongo.Cursor, error)
	DeleteOne(ctx context.Context, filter interface{}, opts ...*options.DeleteOptions) (*mongo.DeleteResult, error)
	UpdateOne(ctx context.Context, filter interface{}, update interface{}, opts ...*options.UpdateOptions) (*mongo.UpdateResult, error)
}

type TaskServiceRepo struct{
	Database_mongo Collectioninter
}

func NewTaskService(Db Collectioninter) *TaskServiceRepo{
	return &TaskServiceRepo{
		Database_mongo : Db,
	}
}

// AddTask adds a new task to the database.
func (this *TaskServiceRepo)SAddTask(task *domain.Task) (*domain.Task, error) {
	store := this.Database_mongo
	data, err := store.InsertOne(context.TODO(), *task)
	id := data.InsertedID.(primitive.ObjectID)
	task.ID = id
	return task, err
}
// GetTask retrieves a task from the database based on the given ID.
func  (this *TaskServiceRepo)SGetTask(id string,user_id string) (*domain.Task, error) {
	var task domain.Task
	ID, _ := primitive.ObjectIDFromHex(id)
	serachIndex:=bson.M{"_id": ID}
	if user_id!=""{
		serachIndex =bson.M{"_id": ID,"user_id":user_id}
	}
	err := this.Database_mongo.FindOne(context.TODO(),serachIndex ).Decode(&task)
	return &task, err
}

// GetTasks retrieves all tasks from the database.
func (this *TaskServiceRepo) SGetTasks(user_id string) ([]domain.Task, error) {
	var tasks []domain.Task
	serachIndex:=bson.M{}
	if user_id!=""{
		serachIndex =bson.M{"user_id":user_id}
	}

	iterDocument, err := this.Database_mongo.Find(context.TODO(),serachIndex)
	for iterDocument.Next(context.TODO()) {
		var task domain.Task
		if err := iterDocument.Decode(&task); err != nil {
			return tasks, err
		}
		tasks = append(tasks, task)
	}
	
	return tasks, err
}

// DeleteTask deletes a task from the database based on the given ID.
func (this *TaskServiceRepo) SDeleteTask(id string) error {
	ID, _ := primitive.ObjectIDFromHex(id)
	check, err := this.Database_mongo.DeleteOne(context.TODO(), bson.M{"_id": ID})
	if check.DeletedCount==0{
		return errors.New("")
	}
	return err
}

// EditTask updates a task in the database based on the given ID.
func  (this *TaskServiceRepo)SEditTask(id string, t *domain.Task) (*domain.Task, error) {
	ID, _:= primitive.ObjectIDFromHex(id)
	update := bson.M{
		"$set": bson.M{
			"title":       t.Title,
			"description": t.Description,
			"due_date":    t.Due_date,
			"status":      t.Status,
		},
	}
	check, err := this.Database_mongo.UpdateOne(context.TODO(), bson.M{"_id": ID}, update)
	t.ID = ID
	if check.MatchedCount==0{
		return t,errors.New("")
	}
	return t, err
}