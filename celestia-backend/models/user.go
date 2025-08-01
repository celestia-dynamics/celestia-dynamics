package models

import(
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type User struct{
	ID primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Name string           `bson:"name" json:"name" binding:"required"`
	Email string          `bson:"email" json:"email" binding:"required,email"`
	Password string       `bson:"password" json:"password,omitempty" binding:"required"`
	CreatedAt time.Time   `bson:"created_at" json:"created_at"`
}