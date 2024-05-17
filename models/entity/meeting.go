package entity

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Meeting struct {
	ID                 primitive.ObjectID   `json:"id" bson:"_id"`
	ProjectId          primitive.ObjectID   `json:"project_id" bson:"project_id"`
	MeetingTitle       string               `json:"meeting_title" bson:"meeting_title"`
	MeetingDescription string               `json:"meeting_description" bson:"meeting_description"`
	MeetingTime        time.Time            `json:"meeting_time" bson:"meeting_time"`
	Users              []primitive.ObjectID `json:"users" bson:"users"`
	CreatedAt          time.Time            `json:"created_at" bson:"created_at"`
	UpdatedAt          time.Time            `json:"updated_at" bson:"updated_at"`
}
