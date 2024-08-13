package controllers

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
	domain "github.com/legend123213/go_togo/Task08/task-manager/Domain"
	"github.com/legend123213/go_togo/Task08/task-manager/mocks"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func Test_AddTask(t *testing.T) {
    gin.SetMode(gin.ReleaseMode)
    mocktask := &mocks.TaskUseCaseint{
        Mock: mock.Mock{
            ExpectedCalls: []*mock.Call{},
            Calls:         []mock.Call{},
        },
    }
    mockcontroller := NewTc(mocktask)
    task := &domain.Task{
        ID:          primitive.ObjectID{},
        Title:       "Test Task",
        Description: "Test Description",
        Due_date:    time.Date(2006, time.January, 2, 15, 4, 5, 0, time.UTC),
        Status:      "Completed",
        UserID:      primitive.ObjectID{0x66, 0xb5, 0xcd, 0xb5, 0xe8, 0x76, 0xe7, 0x9d, 0x4e, 0xe1, 0xb2, 0x45},
    }

    mocktask.On("Create", mock.AnythingOfType("*domain.Task")).Return(task, nil)
    w := httptest.NewRecorder()

    c, _ := gin.CreateTestContext(w)

    c.Request, _ = http.NewRequest(http.MethodPost, "/task", strings.NewReader(`{"user_id":"66b5cdb5e876e79d4ee1b245","title":"Test Task","description":"Test Description","due_date":"2006-01-02T15:04:05Z","status":"Completed"}`))
    c.Request.Header.Set("Content-Type", "application/json")
    mockcontroller.CreateTask(c)
    assert.Equal(t, 201, w.Code)
}

func Test_UpdateTask(t *testing.T) {
    gin.SetMode(gin.ReleaseMode)
    mocktask := &mocks.TaskUseCaseint{
        Mock: mock.Mock{
            ExpectedCalls: []*mock.Call{},
            Calls:         []mock.Call{},
        },
    }
    mockcontroller := NewTc(mocktask)
    taskID := primitive.NewObjectID()
    updatedTask := &domain.Task{
        ID:          taskID,
        Title:       "Updated Task",
        Description: "Updated Description",
        Due_date:    time.Date(2006, time.January, 2, 15, 4, 5, 0, time.UTC),
        Status:      "In Progress",
        UserID:      primitive.NewObjectID(),
    }

    mocktask.On("UpdateTask", taskID.Hex(), mock.AnythingOfType("*domain.Task")).Return(updatedTask, nil)
    w := httptest.NewRecorder()

    c, _ := gin.CreateTestContext(w)
    c.Params = gin.Params{gin.Param{Key: "id", Value: taskID.Hex()}}
    c.Request, _ = http.NewRequest(http.MethodPut, "/task/"+taskID.Hex(), strings.NewReader(`{"title":"Updated Task","description":"Updated Description","due_date":"2006-01-02T15:04:05Z","status":"In Progress"}`))
    c.Request.Header.Set("Content-Type", "application/json")

    mockcontroller.UpdateTask(c)
    assert.Equal(t, 200, w.Code)
}
func Test_GetTask(t *testing.T) {
    gin.SetMode(gin.ReleaseMode)
    mocktask := &mocks.TaskUseCaseint{
        Mock: mock.Mock{
            ExpectedCalls: []*mock.Call{},
            Calls:         []mock.Call{},
        },
    }
    mockcontroller := NewTc(mocktask)
    taskID := primitive.NewObjectID()
    task := &domain.Task{
        ID:          taskID,
        Title:       "Test Task",
        Description: "Test Description",
        Due_date:    time.Date(2006, time.January, 2, 15, 4, 5, 0, time.UTC),
        Status:      "Completed",
        UserID:      primitive.NewObjectID(),
    }

    mocktask.On("FetchTask", taskID.Hex(),mock.Anything).Return(task, nil)
    w := httptest.NewRecorder()

    c, _ := gin.CreateTestContext(w)
    c.Set("isActive",true)
    c.Set("id",primitive.NewObjectID())
    c.Params = gin.Params{gin.Param{Key: "id", Value: taskID.Hex()}}
    c.Request, _ = http.NewRequest(http.MethodGet, "/task/"+taskID.Hex(), nil)

    mockcontroller.GetTask(c)
    assert.Equal(t, 200, w.Code)
}
func Test_SelectTask(t *testing.T) {
    gin.SetMode(gin.ReleaseMode)
    mocktask := &mocks.TaskUseCaseint{
        Mock: mock.Mock{
            ExpectedCalls: []*mock.Call{},
            Calls:         []mock.Call{},
        },
    }
    mockcontroller := NewTc(mocktask)
    taskID := primitive.NewObjectID()
    mocktask.On("RemoveTask", taskID.Hex()).Return(nil)
    w := httptest.NewRecorder()

    c, _ := gin.CreateTestContext(w)
    c.Params = gin.Params{gin.Param{Key: "id", Value: taskID.Hex()}}
    c.Request, _ = http.NewRequest(http.MethodGet, "/task/"+taskID.Hex(), nil)

    mockcontroller.RemoveTask(c)
    assert.Equal(t, 200, w.Code)
}
func Test_FetchAllTasks(t *testing.T) {
    gin.SetMode(gin.ReleaseMode)
    mocktask := &mocks.TaskUseCaseint{
        Mock: mock.Mock{
            ExpectedCalls: []*mock.Call{},
            Calls:         []mock.Call{},
        },
    }
    mockcontroller := NewTc(mocktask)
    tasks := []domain.Task{
        {
            ID:          primitive.NewObjectID(),
            Title:       "Task 1",
            Description: "Description 1",
            Due_date:    time.Date(2006, time.January, 2, 15, 4, 5, 0, time.UTC),
            Status:      "Completed",
            UserID:      primitive.NewObjectID(),
        },
        {
            ID:          primitive.NewObjectID(),
            Title:       "Task 2",
            Description: "Description 2",
            Due_date:    time.Date(2006, time.January, 3, 15, 4, 5, 0, time.UTC),
            Status:      "In Progress",
            UserID:      primitive.NewObjectID(),
        },
    }

    mocktask.On("FetchTasks","").Return(tasks, nil)
    w := httptest.NewRecorder()

    c, _ := gin.CreateTestContext(w)
    c.Request, _ = http.NewRequest(http.MethodGet, "/tasks", nil)
    c.Set("isActive",true)
    c.Set("id",primitive.NewObjectID())
    mockcontroller.GetAllTask(c)
    assert.Equal(t, 200, w.Code)
}



