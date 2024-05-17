package request

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TaskRequest struct {
	ProjectId       primitive.ObjectID   `json:"project_id" bson:"project_id" validate:"required"`
	TaskTitle       string               `json:"task_title" bson:"task_title" validate:"required"`
	TaskDescription string               `json:"task_description" bson:"task_description" validate:"required"`
	TaskStatus      string               `json:"task_status" bson:"task_status" validate:"required,oneof=planning ongoing completed"`
	Users           []primitive.ObjectID `json:"users" bson:"users"`
}

type TaskUpdateRequest struct {
	TaskTitle       string               `json:"task_title" bson:"task_title"`
	TaskDescription string               `json:"task_description" bson:"task_description"`
	TaskStatus      string               `json:"task_status" bson:"task_status" validate:"oneof=planning ongoing completed"`
	Users           []primitive.ObjectID `json:"users" bson:"users"`
}
