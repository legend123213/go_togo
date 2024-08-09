package data

import (
	"context"
	"errors"
	"log"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/legend123213/go_togo/Task06/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

// Generate JWT
func genratetoken(user *models.User,pwd string) (string,error){
	var jwtSecret = []byte(os.Getenv("JWT_SECRET_KEY"))
	// User login logic
	if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(pwd)) != nil {
  return "Invalid username or password",nil
			}

expirationTime := time.Now().Add(24 * 7 * time.Hour)
claims := &Claims{
		ID:       user.ID,
		Username: user.Username,
		IsAdmin:  user.IsAdmin,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
token := jwt.NewWithClaims(jwt.SigningMethodHS256,claims)

jwtToken, err := token.SignedString(jwtSecret)
return jwtToken,err
   
}









type UserServices interface{
	RegisterUser(u *models.User,s *mongo.Database) (string,error)
	EditUser(id string,user *models.User,s *mongo.Database)(*models.User,error)
	GetUser(id string,s *mongo.Database) (*models.User,error)
	DeleteUser(id string,s *mongo.Database) error
	LoginUser(u *models.User,s *mongo.Database) (string,error)
	GetUserByUname(username string,s *mongo.Database)(*models.User,error)
	RoleChanger(id string,s *mongo.Database)(error)
	GetAllUser(s *mongo.Database) *[]models.User
}

type UserServiceRepo struct {}


type Claims struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Username string `json:"username`
	IsAdmin bool    `json:"isadmin"`
	jwt.StandardClaims
}
func NewUserService() *UserServiceRepo{
	return &UserServiceRepo{

	}
}
func (u *UserServiceRepo) RegisterUser(user *models.User,s *mongo.Database) (string,error){
	store := s.Collection("Users")
	pwd:=user.Password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
  		return "Internal server error",err
	}
	user.Password = string(hashedPassword)
	cursor, errFound := store.CountDocuments(context.TODO(), bson.D{})
	if errFound != nil {
		log.Fatal(err)
		return "mongo error",errFound
	}
	user.IsAdmin = true
	if cursor != 0{
		user.IsAdmin = false
	}
	
	data,err_ := store.InsertOne(context.TODO(),*user)
	if err_ != nil {
		return "server error",err_
	}
	id := data.InsertedID.(primitive.ObjectID)
	user.ID= id
	validUser,err := u.GetUser(id.Hex(),s)
	if err!=nil{
		return "user not found",err
	}
	token,err := genratetoken(validUser,pwd)
	return token,err
}
func (u *UserServiceRepo) EditUser(id string,user *models.User,s *mongo.Database) (*models.User,error){
	ID, err_:= primitive.ObjectIDFromHex(id)
	log.Println(err_)
	update := bson.M{
		"$set": bson.M{
			"username":   user.Username,
		},
	}
	check, err := s.Collection("Tasks").UpdateOne(context.TODO(), bson.M{"_id": ID}, update)
	user.ID = ID
	if check.MatchedCount==0{
		return user,errors.New("")
	}
	return user, err
}
func (u *UserServiceRepo) GetUser(id string,s *mongo.Database) (*models.User,error){
	var user models.User
	ID, _ := primitive.ObjectIDFromHex(id)
	err := s.Collection("Users").FindOne(context.TODO(), bson.M{"_id": ID}).Decode(&user)
	return &user, err
}
func (u *UserServiceRepo) GetUserByUname(username string,s *mongo.Database)(*models.User,error){
	var user models.User
	err := s.Collection("Users").FindOne(context.TODO(), bson.M{"username": username}).Decode(&user)
	return &user, err

}
func (u *UserServiceRepo) DeleteUser(id string,s *mongo.Database) error{
	
	log.Println(id,s,id)
	return nil
}
func (u *UserServiceRepo) LoginUser(user *models.User,s *mongo.Database) (string,error){
	pwd:=user.Password
	User,err := u.GetUserByUname(user.Username,s)
	if err !=nil{
		return "user not found",errors.New("")
	}

	return genratetoken(User,pwd)
}
func (u *UserServiceRepo) RoleChanger(id string,s *mongo.Database)(error){
	User,err := u.GetUser(id,s)
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
			"role": User.IsAdmin,
		},
	}
	check, err := s.Collection("Users").UpdateOne(context.TODO(), bson.M{"_id":User.ID}, update)
	if err != nil {
		log.Println(err)
		return errors.New("")
	}
	if check.MatchedCount==0{
		return errors.New("")
	}
	return nil
}

func(u *UserServiceRepo) GetAllUser(s *mongo.Database) *[]models.User{
	var users []models.User
	store,_ := s.Collection("Users").Find(context.TODO(),bson.D{})
	for store.Next(context.TODO()){
		var user models.User
		store.Decode(&user)
		users=append(users,user)
	}
	
	return &users
}