package handlers

import (
	"github.com/gofiber/fiber/v2"
	"go-project-management/database"
	"go-project-management/models/entity"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetAllMeetings(ctx *fiber.Ctx) error {
	collection := database.GetDBCollection("meetings")

	var meetings []entity.Meeting

	cursor, err := collection.Find(ctx.Context(), primitive.M{})
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to get all meetings",
			"error":   err.Error(),
		})
	}

	if cursor.RemainingBatchLength() == 0 {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "No meetings found",
		})
	}

	err = cursor.All(ctx.Context(), &meetings)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to get all meetings",
			"error":   err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Success get all meetings",
		"data":    meetings,
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
