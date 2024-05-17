package entity

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Project struct {
	ID          primitive.ObjectID   `json:"id" bson:"_id"`
	Title       string               `json:"title" bson:"title"`
	Description string               `json:"description" bson:"description"`
	Status      string               `json:"status" bson:"status"` // planning, ongoing, completed
	AdminUserId primitive.ObjectID   `json:"admin_user_id" bson:"admin_user_id"`
	StartDate   time.Time            `json:"start_date" bson:"start_date"`
	EndDate     time.Time            `json:"end_date" bson:"end_date"`
	Users       []primitive.ObjectID `json:"users" bson:"users"`
	CreatedAt   time.Time            `json:"created_at" bson:"created_at"`
	UpdatedAt   time.Time            `json:"updated_at" bson:"updated_at"`
}
