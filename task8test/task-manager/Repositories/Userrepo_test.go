package repositories

import (
	"fmt"
	"log"
	"testing"

	domain "github.com/legend123213/go_togo/Task08/task-manager/Domain"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/integration/mtest"
	"golang.org/x/crypto/bcrypt"
)

func Test_IsUsernameUnique(t *testing.T) {
    mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
    mt.Run("username exists", func(mt *mtest.T) {
        userCollection := mt.Coll
        repo := NewUserService(userCollection)
        username := "existing_user"

        mt.AddMockResponses(mtest.CreateCursorResponse(1, "foo.bar", mtest.FirstBatch, bson.D{
            {"username", username},
        }))

        err := repo.IsUsernameUnique(username)
        assert.NotNil(t, err)
        assert.Equal(t, fmt.Sprintf("username '%s' already exists", username), err.Error())
    })

    mt.Run("username does not exist", func(mt *mtest.T) {
        userCollection := mt.Coll
        repo := NewUserService(userCollection)
        username := "new_user"

        mt.AddMockResponses(mtest.CreateCursorResponse(0, "foo.bar", mtest.FirstBatch))

        err := repo.IsUsernameUnique(username)
        assert.Nil(t, err)
    })
}

func Test_RegisterUser(t *testing.T) {
    mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
    mt.Run("success", func(mt *mtest.T) {
        userCollection := mt.Coll
        repo := NewUserService(userCollection)
        newUser := &domain.User{
            Username: "new_user",
            Password: "password123",
        }

        mt.AddMockResponses(mtest.CreateSuccessResponse())

        token, err := repo.RegisterUser(newUser)
		  log.Println(token)
        assert.Nil(t, err)
        assert.NotEmpty(t, token)
    })
}

func Test_EditUser(t *testing.T) {
    mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
    mt.Run("success", func(mt *mtest.T) {
        userCollection := mt.Coll
        repo := NewUserService(userCollection)
        id := primitive.NewObjectID()
        updatedUser := &domain.User{
            Username: "updated_user",
        }

        mt.AddMockResponses(bson.D{
            {"ok", 1},
            {"value", bson.D{
                {"username", updatedUser.Username},
            }},
        })

        user, err := repo.EditUser(id.Hex(), updatedUser)
        assert.Nil(t, err)
        assert.Equal(t, updatedUser.Username, user.Username)
    })
}

func Test_GetUser(t *testing.T) {
    mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
    mt.Run("success", func(mt *mtest.T) {
        userCollection := mt.Coll
        repo := NewUserService(userCollection)
        id := primitive.NewObjectID()
        expectedUser := &domain.User{
            ID:       id,
            Username: "test_user",
        }

        mt.AddMockResponses(mtest.CreateCursorResponse(1, "foo.bar", mtest.FirstBatch, bson.D{
            {"_id", expectedUser.ID},
            {"username", expectedUser.Username},
        }))

        user, err := repo.GetUser(id.Hex())
        assert.Nil(t, err)
        assert.Equal(t, expectedUser.Username, user.Username)
    })
}

func Test_GetUserByUname(t *testing.T) {
    mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
    mt.Run("success", func(mt *mtest.T) {
        userCollection := mt.Coll
        repo := NewUserService(userCollection)
        username := "test_user"
        expectedUser := &domain.User{
            Username: username,
        }

        mt.AddMockResponses(mtest.CreateCursorResponse(1, "foo.bar", mtest.FirstBatch, bson.D{
            {"username", expectedUser.Username},
        }))

        user, err := repo.GetUserByUname(username)
        assert.Nil(t, err)
        assert.Equal(t, expectedUser.Username, user.Username)
    })
}

func Test_DeleteUser(t *testing.T) {
    mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
    mt.Run("success", func(mt *mtest.T) {
        userCollection := mt.Coll
        repo := NewUserService(userCollection)
        id := primitive.NewObjectID()

        mt.AddMockResponses(mtest.CreateSuccessResponse())

        err := repo.DeleteUser(id.Hex())
        assert.Nil(t, err)
    })
}

func Test_LoginUser(t *testing.T) {
    mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
    mt.Run("success", func(mt *mtest.T) {
        userCollection := mt.Coll
        repo := NewUserService(userCollection)
        username := "test_user"
        password := "password123"
        hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
        expectedUser := &domain.User{
            Username: username,
            Password: string(hashedPassword),
        }

        mt.AddMockResponses(mtest.CreateCursorResponse(1, "foo.bar", mtest.FirstBatch, bson.D{
            {"username", expectedUser.Username},
            {"password", expectedUser.Password},
        }))

        token, err := repo.LoginUser(&domain.User{Username: username, Password: password})
        assert.Nil(t, err)
        assert.NotEmpty(t, token)
    })
}