package entity

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Task struct {
	ID              primitive.ObjectID   `json:"id" bson:"_id"`
	ProjectId       primitive.ObjectID   `json:"project_id" bson:"project_id"`
	TaskTitle       string               `json:"task_title" bson:"task_title"`
	TaskDescription string               `json:"task_description" bson:"task_description"`
	TaskStatus      string               `json:"task_status" bson:"task_status"` // planning, ongoing, completed
	Users           []primitive.ObjectID `json:"users" bson:"users"`
	CreatedAt       time.Time            `json:"created_at" bson:"created_at"`
	UpdatedAt       time.Time            `json:"updated_at" bson:"updated_at"`
}
