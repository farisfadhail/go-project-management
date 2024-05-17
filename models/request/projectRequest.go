package request

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type ProjectRequest struct {
	Title       string               `json:"title" bson:"title" validate:"required"`
	Description string               `json:"description" bson:"description" validate:"required"`
	Status      string               `json:"status" bson:"status" validate:"required,oneof=planning ongoing completed"`
	AdminUserId primitive.ObjectID   `json:"admin_user_id" bson:"admin_user_id" validate:"required"`
	StartDate   time.Time            `json:"start_date" bson:"start_date" validate:"required"`
	EndDate     time.Time            `json:"end_date" bson:"end_date" validate:"required"`
	Users       []primitive.ObjectID `json:"users" bson:"users" validate:"required"`
}

type ProjectUpdateRequest struct {
	Title       string               `json:"title" bson:"title"`
	Description string               `json:"description" bson:"description"`
	Status      string               `json:"status" bson:"status" validate:"oneof=planning ongoing completed"`
	EndDate     time.Time            `json:"end_date" bson:"end_date"`
	Users       []primitive.ObjectID `json:"users" bson:"users"`
}
