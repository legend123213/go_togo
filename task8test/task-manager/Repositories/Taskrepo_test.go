package repositories

import (
	"context"
	"log"
	"testing"
	"time"

	domain "github.com/legend123213/go_togo/Task08/task-manager/Domain"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// MockMongoCollection is a mock of the mongo.Collection
type MockMongoCollection struct {
    mock.Mock
}

func (m *MockMongoCollection) InsertOne(ctx context.Context, document interface{}, opts ...*options.InsertOneOptions) (*mongo.InsertOneResult, error) {
    args := m.Called(ctx, document, opts)
    return args.Get(0).(*mongo.InsertOneResult), args.Error(1)
}

func (m *MockMongoCollection) FindOne(ctx context.Context, filter interface{}, opts ...*options.FindOneOptions) *mongo.SingleResult {
    args := m.Called(ctx, filter, opts)
    return args.Get(0).(*mongo.SingleResult)
}

func (m *MockMongoCollection) Find(ctx context.Context, filter interface{}, opts ...*options.FindOptions) (*mongo.Cursor, error) {
    args := m.Called(ctx, filter, opts)
    return args.Get(0).(*mongo.Cursor), args.Error(1)
}

func (m *MockMongoCollection) DeleteOne(ctx context.Context, filter interface{}, opts ...*options.DeleteOptions) (*mongo.DeleteResult, error) {
    args := m.Called(ctx, filter, opts)
    return args.Get(0).(*mongo.DeleteResult), args.Error(1)
}

func (m *MockMongoCollection) UpdateOne(ctx context.Context, filter interface{}, update interface{}, opts ...*options.UpdateOptions) (*mongo.UpdateResult, error) {
    args := m.Called(ctx, filter, update, opts)
    return args.Get(0).(*mongo.UpdateResult), args.Error(1)
}


func TestSAddTask(t *testing.T) {
    mockcol := &MockMongoCollection{}
    taskService:= NewTaskService(mockcol)

    task := &domain.Task{
        Title:       "Test Task",
        Description: "Test Description",
        Due_date:    time.Now(),
        Status:      "Pending",
    }

    mockcol.On("InsertOne", mock.Anything, *task, mock.Anything).Return(&mongo.InsertOneResult{
        InsertedID: primitive.NewObjectID(),
    }, nil)

    result, err := taskService.SAddTask(task)
    assert.NoError(t, err)
    assert.NotNil(t, result.ID)
    log.Println(result,task)
    
    assert.Equal(t,result,task)
}

