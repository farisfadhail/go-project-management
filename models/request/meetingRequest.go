package request

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type MeetingRequest struct {
	ProjectId          primitive.ObjectID   `json:"project_id" bson:"project_id" validate:"required"`
	MeetingTitle       string               `json:"meeting_title" bson:"meeting_title" validate:"required"`
	MeetingDescription string               `json:"meeting_description" bson:"meeting_description" validate:"required"`
	MeetingTime        time.Time            `json:"meeting_time" bson:"meeting_time" validate:"required"`
	Users              []primitive.ObjectID `json:"users" bson:"users" validate:"required"`
}

type MeetingUpdateRequest struct {
	MeetingTitle       string               `json:"meeting_title" bson:"meeting_title"`
	MeetingDescription string               `json:"meeting_description" bson:"meeting_description"`
	MeetingTime        time.Time            `json:"meeting_time" bson:"meeting_time" `
	Users              []primitive.ObjectID `json:"users" bson:"users"`
}
