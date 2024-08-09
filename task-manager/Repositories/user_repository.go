package repositories

import (
	"context"
	"errors"
	"fmt"
	"log"

	domain "github.com/legend123213/go_togo/Task07/task-manager/Domain"
	Infrastructure "github.com/legend123213/go_togo/Task07/task-manager/Infrastructure"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

// UserServices is an interface that defines the methods for user management.
type UserServices interface {
	RegisterUser(u *domain.User) (string, error)
	EditUser(id string, user *domain.User) (*domain.User, error)
	GetUser(id string) (*domain.User, error)
	DeleteUser(id string) error
	LoginUser(u *domain.User) (string, error)
	GetUserByUname(username string) (*domain.User, error)
	RoleChanger(id string) error
	GetAllUser() *[]domain.User
	IsUsernameUnique(username string) error
}

// UserServiceRepo is a struct that implements the UserServices interface.
type UserServiceRepo struct {
	Database_mongo *mongo.Database
}

// NewUserService creates a new instance of UserServiceRepo.
func NewUserService(Db *mongo.Database) *UserServiceRepo {
	return &UserServiceRepo{
		Database_mongo: Db,
	}
}

// IsUsernameUnique checks if the given username is unique.
func (u *UserServiceRepo) IsUsernameUnique(username string) error {
	var existingUser domain.User
	store := u.Database_mongo.Collection("Users")
	err := store.FindOne(context.TODO(), bson.M{"username": username}).Decode(&existingUser)
	if err == nil {
		return fmt.Errorf("username '%s' already exists", username)
	}
	if err == mongo.ErrNoDocuments {
		return nil
	}

	return err
}

// RegisterUser registers a new user.
func (u *UserServiceRepo) RegisterUser(user *domain.User) (string, error) {
	store := u.Database_mongo.Collection("Users")
	pwd := user.Password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return "Internal server error", err
	}
	user.Password = string(hashedPassword)
	cursor, errFound := store.CountDocuments(context.TODO(), bson.D{})
	if errFound != nil {
		log.Fatal(err)
		return "mongo error", errFound
	}
	user.IsAdmin = true
	if cursor != 0 {
		user.IsAdmin = false
	}

	data, err_ := store.InsertOne(context.TODO(), *user)
	if err_ != nil {
		return "server error", err_
	}
	id := data.InsertedID.(primitive.ObjectID)
	user.ID = id
	validUser, err := u.GetUser(id.Hex())
	if err != nil {
		return "user not found", err
	}
	token, err := Infrastructure.Genratetoken(validUser, pwd)
	return token, err
}

// EditUser updates the details of a user.
func (u *UserServiceRepo) EditUser(id string, user *domain.User) (*domain.User, error) {
	ID, err_ := primitive.ObjectIDFromHex(id)
	log.Println(err_)
	update := bson.M{
		"$set": bson.M{
			"username": user.Username,
		},
	}
	check, err := u.Database_mongo.Collection("Tasks").UpdateOne(context.TODO(), bson.M{"_id": ID}, update)
	user.ID = ID
	if check.MatchedCount == 0 {
		return user, errors.New("")
	}
	return user, err
}

// GetUser retrieves a user by ID.
func (u *UserServiceRepo) GetUser(id string) (*domain.User, error) {
	var user domain.User
	ID, _ := primitive.ObjectIDFromHex(id)
	err := u.Database_mongo.Collection("Users").FindOne(context.TODO(), bson.M{"_id": ID}).Decode(&user)
	return &user, err
}

// GetUserByUname retrieves a user by username.
func (u *UserServiceRepo) GetUserByUname(username string) (*domain.User, error) {
	var user domain.User
	err := u.Database_mongo.Collection("Users").FindOne(context.TODO(), bson.M{"username": username}).Decode(&user)
	return &user, err
}

// DeleteUser deletes a user by ID.
func (u *UserServiceRepo) DeleteUser(id string) error {
	ID, _ := primitive.ObjectIDFromHex(id)
	_, err := u.Database_mongo.Collection("Users").DeleteOne(context.TODO(), bson.M{"_id": ID})
	if err != nil {
		return err
	}

	return err
}

// LoginUser performs user login and returns a JWT token.
func (u *UserServiceRepo) LoginUser(user *domain.User) (string, error) {
	pwd := user.Password
	User, err := u.GetUserByUname(user.Username)
	if err != nil {
		return "user not found", errors.New("")
	}

	return Infrastructure.Genratetoken(User, pwd)
}

// RoleChanger changes the role of a user to admin.
func (u *UserServiceRepo) RoleChanger(id string) error {
	User, err := u.GetUser(id)
	if err != nil {
		return errors.New("user not found")
	}
	if User.IsAdmin {
		return errors.New("user already admin")
	}
	User.IsAdmin = true
	update := bson.M{
		"$set": bson.M{
			"username": User.Username,
			"role":     User.IsAdmin,
		},
	}
	check, err := u.Database_mongo.Collection("Users").UpdateOne(context.TODO(), bson.M{"_id": User.ID}, update)
	if err != nil {
		log.Println(err)
		return errors.New("")
	}
	if check.MatchedCount == 0 {
		return errors.New("")
	}
	return nil
}

// GetAllUser retrieves all users.
func (u *UserServiceRepo) GetAllUser() *[]domain.User {
	var users []domain.User
	store, _ := u.Database_mongo.Collection("Users").Find(context.TODO(), bson.D{})
	for store.Next(context.TODO()) {
		var user domain.User
		store.Decode(&user)
		users = append(users, user)
	}

	return &users
}
