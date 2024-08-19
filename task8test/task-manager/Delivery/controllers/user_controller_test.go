package controllers

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
	domain "github.com/legend123213/go_togo/Task08/task-manager/Domain"
	"github.com/legend123213/go_togo/Task08/task-manager/mocks"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func Test_CreateUser(t *testing.T) {
    gin.SetMode(gin.ReleaseMode)
    mockUserUsecase := &mocks.UserUsecaseInt{
        Mock: mock.Mock{
            ExpectedCalls: []*mock.Call{},
            Calls:         []mock.Call{},
        },
    }
    mockController := NewUc(mockUserUsecase)
    user := &domain.User{
        Username: "testuser",
        Password: "password",
    }

    mockUserUsecase.On("IsUsernameUnique", user.Username).Return(nil)
    mockUserUsecase.On("Register", mock.AnythingOfType("*domain.User")).Return("token123", nil)
    w := httptest.NewRecorder()

    c, _ := gin.CreateTestContext(w)
    c.Request, _ = http.NewRequest(http.MethodPost, "/user", strings.NewReader(`{"username":"testuser","password":"password"}`))
    c.Request.Header.Set("Content-Type", "application/json")

    mockController.CreateUser(c)
    assert.Equal(t, 201, w.Code)
}

func Test_EditUser(t *testing.T) {
    gin.SetMode(gin.ReleaseMode)
    mockUserUsecase := &mocks.UserUsecaseInt{
        Mock: mock.Mock{
            ExpectedCalls: []*mock.Call{},
            Calls:         []mock.Call{},
        },
    }
    mockController := NewUc(mockUserUsecase)
    userID := primitive.NewObjectID().Hex()
    updatedUser := &domain.User{
        Username: "updateduser",
    }

    mockUserUsecase.On("Edit", userID, mock.AnythingOfType("*domain.User")).Return(updatedUser, nil)
    w := httptest.NewRecorder()

    c, _ := gin.CreateTestContext(w)
    c.Params = gin.Params{gin.Param{Key: "id", Value: userID}}
    c.Request, _ = http.NewRequest(http.MethodPut, "/user/"+userID, strings.NewReader(`{"username":"updateduser"}`))
    c.Request.Header.Set("Content-Type", "application/json")

    mockController.UpdateUser(c)
    assert.Equal(t, 200, w.Code)
}

func Test_GetUser(t *testing.T) {
    gin.SetMode(gin.ReleaseMode)
    mockUserUsecase := &mocks.UserUsecaseInt{
        Mock: mock.Mock{
            ExpectedCalls: []*mock.Call{},
            Calls:         []mock.Call{},
        },
    }
    mockController := NewUc(mockUserUsecase)
    userID := "12345"
    user := &domain.User{
        Username: "testuser",
        Password: "password",
        IsAdmin: true,
    }

    mockUserUsecase.On("Fetch", userID).Return(user, nil)
    w := httptest.NewRecorder()

    c, _ := gin.CreateTestContext(w)
    c.Params = gin.Params{gin.Param{Key: "id", Value: userID}}
    c.Request, _ = http.NewRequest(http.MethodGet, "/user/"+userID, nil)

    mockController.GetUser(c)
    assert.Equal(t, 200, w.Code)
}

func Test_DeleteUser(t *testing.T) {
    gin.SetMode(gin.ReleaseMode)
    mockUserUsecase := &mocks.UserUsecaseInt{
        Mock: mock.Mock{
            ExpectedCalls: []*mock.Call{},
            Calls:         []mock.Call{},
        },
    }
    mockController := NewUc(mockUserUsecase)
    userID := primitive.NewObjectID().Hex()

    mockUserUsecase.On("Delete", userID).Return(nil)
    w := httptest.NewRecorder()

    c, _ := gin.CreateTestContext(w)
    c.Params = gin.Params{gin.Param{Key: "id", Value: userID}}
    c.Request, _ = http.NewRequest(http.MethodDelete, "/user/"+userID, nil)

    mockController.RemoveUser(c)
    assert.Equal(t, 200, w.Code)
}

func Test_LoginUser(t *testing.T) {
    gin.SetMode(gin.ReleaseMode)
    mockUserUsecase := &mocks.UserUsecaseInt{
        Mock: mock.Mock{
            ExpectedCalls: []*mock.Call{},
            Calls:         []mock.Call{},
        },
    }
    mockController := NewUc(mockUserUsecase)

    mockUserUsecase.On("Login", mock.AnythingOfType("*domain.User")).Return("token123", nil)
    w := httptest.NewRecorder()

    c, _ := gin.CreateTestContext(w)
    c.Request, _ = http.NewRequest(http.MethodPost, "/login", strings.NewReader(`{"username":"testuser","password":"password"}`))
    c.Request.Header.Set("Content-Type", "application/json")

    mockController.LogUser(c)
    assert.Equal(t, 200, w.Code)
}