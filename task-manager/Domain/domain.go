package domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Task represents a task in the system.
type Task struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Title       string             `json:"title" bson:"title"`
	Description string             `json:"description" bson:"description"`
	Due_date     time.Time          `json:"due_date" bson:"due_date"`
	Status      string             `json:"status" bson:"status"`
	UserID      primitive.ObjectID  `bson:"user_id," json:"user_id"`
}



// User represents a task in the system.
type User struct{
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Username string `bson:"username" json:"username"`
	IsAdmin bool `bson:"role" json:"role"`
	Password string `bson:"pwd" json:"-"`
}
