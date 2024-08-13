package usecases

import (
	"testing"
	"time"

	domain "github.com/legend123213/go_togo/Task08/task-manager/Domain"
	"github.com/legend123213/go_togo/Task08/task-manager/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func Test_addtask(t *testing.T) {
    mocktaskrepo := &mocks.TaskInterface{
        Mock: mock.Mock{
            ExpectedCalls: []*mock.Call{},
            Calls:         []mock.Call{},
        },
    }
    mockusecase := NewTaskUsecase(mocktaskrepo)
    taskID := primitive.NewObjectID()
    task := &domain.Task{
        ID:          taskID,
        Title:       "Test Task",
        Description: "Test Description",
        Due_date:    time.Date(2006, time.January, 2, 15, 4, 5, 0, time.UTC),
        Status:      "Completed",
        UserID:      primitive.NewObjectID(),
    }
    mocktaskrepo.On("SAddTask", task).Return(task, nil)
    res, err := mockusecase.Create(task)
    assert.Nil(t, err)
    assert.Equal(t, res, task)
}

func Test_FetchTasks(t *testing.T) {
    mocktaskrepo := &mocks.TaskInterface{
        Mock: mock.Mock{
            ExpectedCalls: []*mock.Call{},
            Calls:         []mock.Call{},
        },
    }
    mockusecase := NewTaskUsecase(mocktaskrepo)
    userID := primitive.NewObjectID().Hex()
    tasks := []domain.Task{
        {
            ID:          primitive.NewObjectID(),
            Title:       "Task 1",
            Description: "Description 1",
            Due_date:    time.Now(),
            Status:      "Pending",
            UserID:      primitive.NewObjectID(),
        },
        {
            ID:          primitive.NewObjectID(),
            Title:       "Task 2",
            Description: "Description 2",
            Due_date:    time.Now(),
            Status:      "Completed",
            UserID:      primitive.NewObjectID(),
        },
    }
    mocktaskrepo.On("SGetTasks", userID).Return(tasks, nil)
    res, err := mockusecase.FetchTasks(userID)
    assert.Nil(t, err)
    assert.Equal(t, res, tasks)
}

func Test_FetchTask(t *testing.T) {
    mocktaskrepo := &mocks.TaskInterface{
        Mock: mock.Mock{
            ExpectedCalls: []*mock.Call{},
            Calls:         []mock.Call{},
        },
    }
    mockusecase := NewTaskUsecase(mocktaskrepo)
    taskID := primitive.NewObjectID().Hex()
    username := "testuser"
    task := &domain.Task{
        ID:          primitive.NewObjectID(),
        Title:       "Test Task",
        Description: "Test Description",
        Due_date:    time.Now(),
        Status:      "Pending",
        UserID:      primitive.NewObjectID(),
    }
    mocktaskrepo.On("SGetTask", taskID, username).Return(task, nil)
    res, err := mockusecase.FetchTask(taskID, username)
    assert.Nil(t, err)
    assert.Equal(t, res, task)
}

func Test_RemoveTask(t *testing.T) {
    mocktaskrepo := &mocks.TaskInterface{
        Mock: mock.Mock{
            ExpectedCalls: []*mock.Call{},
            Calls:         []mock.Call{},
        },
    }
    mockusecase := NewTaskUsecase(mocktaskrepo)
    taskID := primitive.NewObjectID().Hex()
    mocktaskrepo.On("SDeleteTask", taskID).Return(nil)
    err := mockusecase.RemoveTask(taskID)
    assert.Nil(t, err)
}

func Test_UpdateTask(t *testing.T) {
    mocktaskrepo := &mocks.TaskInterface{
        Mock: mock.Mock{
            ExpectedCalls: []*mock.Call{},
            Calls:         []mock.Call{},
        },
    }
    mockusecase := NewTaskUsecase(mocktaskrepo)
    taskID := primitive.NewObjectID().Hex()
    task := &domain.Task{
        ID:          primitive.NewObjectID(),
        Title:       "Updated Task",
        Description: "Updated Description",
        Due_date:    time.Now(),
        Status:      "In Progress",
        UserID:      primitive.NewObjectID(),
    }
    mocktaskrepo.On("SEditTask", taskID, task).Return(task, nil)
    res, err := mockusecase.UpdateTask(taskID, task)
    assert.Nil(t, err)
    assert.Equal(t, res, task)
}