package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"go-project-management/database"
	"go-project-management/models/entity"
	"go-project-management/models/request"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

func CreateMeeting(ctx *fiber.Ctx) error {
	meetingRequest := new(request.MeetingRequest)
	err := ctx.BodyParser(meetingRequest)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Failed to parse meeting request",
			"error":   err.Error(),
		})
	}

	findOne := database.GetDBCollection("projects").FindOne(ctx.Context(), primitive.M{"_id": meetingRequest.ProjectId})
	if findOne.Err() != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Project not found",
			"error":   findOne.Err().Error(),
		})
	}

	newMeeting := entity.Meeting{
		ID:                 primitive.NewObjectID(),
		ProjectId:          meetingRequest.ProjectId,
		MeetingTitle:       meetingRequest.MeetingTitle,
		MeetingDescription: meetingRequest.MeetingDescription,
		MeetingTime:        meetingRequest.MeetingTime,
		Users:              meetingRequest.Users,
		CreatedAt:          time.Now(),
		UpdatedAt:          time.Now(),
	}

	collection := database.GetDBCollection("meetings")

	result, err := collection.InsertOne(ctx.Context(), newMeeting)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to create meeting",
			"error":   err.Error(),
		})
	}

	log.Info("Meeting created: ", result.InsertedID)

	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Success create meeting",
		"data":    newMeeting,
	})
}

func GetMeetingById(ctx *fiber.Ctx) error {
	collection := database.GetDBCollection("meetings")

	meetingId := ctx.Params("meetingId")

	objectID, err := primitive.ObjectIDFromHex(meetingId)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid meeting ID",
			"error":   err.Error(),
		})
	}

	var meeting entity.Meeting

	err = collection.FindOne(ctx.Context(), primitive.M{"_id": objectID}).Decode(&meeting)
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Meeting not found",
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Success get meeting",
		"data":    meeting,
	})
}

func UpdateMeeting(ctx *fiber.Ctx) error {
	updateRequest := new(request.MeetingUpdateRequest)
	err := ctx.BodyParser(updateRequest)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Failed to parse meeting update request",
			"error":   err.Error(),
		})
	}

	collection := database.GetDBCollection("meetings")

	meetingId := ctx.Params("meetingId")

	objectID, err := primitive.ObjectIDFromHex(meetingId)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid meeting ID",
			"error":   err.Error(),
		})
	}

	var meeting entity.Meeting

	err = collection.FindOne(ctx.Context(), primitive.M{"_id": objectID}).Decode(&meeting)
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Meeting not found",
		})
	}

	meeting.MeetingTitle = updateRequest.MeetingTitle
	meeting.MeetingDescription = updateRequest.MeetingDescription
	meeting.MeetingTime = updateRequest.MeetingTime
	meeting.Users = updateRequest.Users
	meeting.UpdatedAt = time.Now()

	_, err = collection.UpdateOne(ctx.Context(), primitive.M{"_id": objectID}, primitive.M{"$set": meeting})
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to update meeting",
			"error":   err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Success update meeting",
		"data":    meeting,
	})
}

func DeleteMeeting(ctx *fiber.Ctx) error {
	collection := database.GetDBCollection("meetings")

	meetingId := ctx.Params("meetingId")

	objectID, err := primitive.ObjectIDFromHex(meetingId)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid meeting ID",
			"error":   err.Error(),
		})
	}

	deleteOne, err := collection.DeleteOne(ctx.Context(), primitive.M{"_id": objectID})
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to delete meeting",
			"error":   err.Error(),
		})
	}

	if deleteOne.DeletedCount == 0 {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Meeting not found",
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Success delete meeting",
	})
}
