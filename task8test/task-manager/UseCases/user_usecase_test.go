package usecases

import (
	"testing"

	domain "github.com/legend123213/go_togo/Task08/task-manager/Domain"
	"github.com/legend123213/go_togo/Task08/task-manager/mocks"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestDelete(t *testing.T) {
    mockUserrepo := new(mocks.UserServices)
	 mockUserusecase:= NewuserUsecase(mockUserrepo)
	 id := primitive.NewObjectID().Hex()
primitive.NewObjectID().Hex()
    mockUserrepo.On("DeleteUser", id).Return(nil)

    err := mockUserusecase.Delete(id)
    if err != nil {
        t.Errorf("expected no error, but got %v", err)
    }

    assert.Nil(t,err)
}

func TestEdit(t *testing.T) {
    mockUserrepo := new(mocks.UserServices)
	 mockUserusecase:= NewuserUsecase(mockUserrepo)
    user := &domain.User{ID: primitive.NewObjectID(), Username: "testuser"}
	 id := primitive.NewObjectID().Hex()
    mockUserrepo.On("EditUser", id, user).Return(user, nil)

    editedUser, err := mockUserusecase.Edit(id, user)
    if err != nil {
        t.Errorf("expected no error, but got %v", err)
    }
    if editedUser.Username != "testuser" {
        t.Errorf("expected username to be 'testuser', but got %v", editedUser.Username)
    }

}
func TestFetchAll(t *testing.T) {
	mockUserrepo := new(mocks.UserServices)
	mockUserusecase := NewuserUsecase(mockUserrepo)
	users := []domain.User{
		{ID: primitive.NewObjectID(), Username: "user1"},
		{ID: primitive.NewObjectID(), Username: "user2"},
		{ID: primitive.NewObjectID(), Username: "user3"},
	}
	mockUserrepo.On("GetAllUser").Return(&users, nil)

	fetchedUsers := mockUserusecase.FetchAllUser()
	assert.Equal(t, len(users), len(*fetchedUsers))
}

func TestFetch(t *testing.T) {
	mockUserrepo := new(mocks.UserServices)
	mockUserusecase := NewuserUsecase(mockUserrepo)
	user := &domain.User{ID: primitive.NewObjectID(), Username: "testuser"}
	id := primitive.NewObjectID().Hex()
	mockUserrepo.On("GetUser", id).Return(user, nil)

	fetchedUser, err := mockUserusecase.Fetch(id)
	assert.Nil(t, err)
	assert.Equal(t, "testuser", fetchedUser.Username)

	mockUserrepo.AssertExpectations(t)
}

func TestFetchUserByUname(t *testing.T) {
	mockUserrepo := new(mocks.UserServices)
	mockUserusecase := NewuserUsecase(mockUserrepo)
	user := &domain.User{ID: primitive.NewObjectID(), Username: "testuser"}
	mockUserrepo.On("GetUserByUname", "testuser").Return(user, nil)

	fetchedUser, err := mockUserusecase.FetchUserByUname("testuser")
	assert.Nil(t, err)
	assert.Equal(t, "testuser", fetchedUser.Username)

	mockUserrepo.AssertExpectations(t)
}

func TestIsUsernameUnique(t *testing.T) {
	mockUserrepo := new(mocks.UserServices)
	mockUserusecase := NewuserUsecase(mockUserrepo)
	mockUserrepo.On("IsUsernameUnique", "testuser").Return(nil)

	err := mockUserusecase.IsUsernameUnique("testuser")
	if err != nil {
		t.Errorf("expected no error, but got %v", err)
	}

	mockUserrepo.AssertExpectations(t)
}

func TestLogin(t *testing.T) {
	mockUserrepo := new(mocks.UserServices)
	mockUserusecase := NewuserUsecase(mockUserrepo)
	user := &domain.User{ID: primitive.NewObjectID(), Username: "testuser"}
	mockUserrepo.On("LoginUser", user).Return("token", nil)

	token, err := mockUserusecase.Login(user)
	assert.Nil(t, err)
	assert.Equal(t, "token", token)

	mockUserrepo.AssertExpectations(t)
}

func TestRegister(t *testing.T) {
	mockUserrepo := new(mocks.UserServices)
	mockUserusecase := NewuserUsecase(mockUserrepo)
	user := &domain.User{ID: primitive.NewObjectID(), Username: "testuser"}
	mockUserrepo.On("RegisterUser", user).Return("token", nil)

	token, err := mockUserusecase.Register(user)
	assert.Nil(t, err)
	assert.Equal(t, "token", token)

	mockUserrepo.AssertExpectations(t)
}

func TestRoleChanger(t *testing.T) {
	mockUserrepo := new(mocks.UserServices)
	mockUserusecase := NewuserUsecase(mockUserrepo)
	id := primitive.NewObjectID().Hex()
	mockUserrepo.On("RoleChanger", id).Return(nil)

	err := mockUserusecase.RoleChanger(id)
	assert.Nil(t, err)
	mockUserrepo.AssertExpectations(t)
}