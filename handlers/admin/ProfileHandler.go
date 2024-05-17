package handlers

import (
	"github.com/gofiber/fiber/v2"
	"go-project-management/database"
	"go-project-management/models/entity"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetAllProfiles(ctx *fiber.Ctx) error {
	collection := database.GetDBCollection("profiles")

	var profiles []entity.Profile

	cursor, err := collection.Find(ctx.Context(), primitive.M{})
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to get all profiles",
			"error":   err.Error(),
		})
	}

	if cursor.RemainingBatchLength() == 0 {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "No meetings found",
		})
	}

	err = cursor.All(ctx.Context(), &profiles)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to get all profiles",
			"error":   err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Success get all profiles",
		"data":    profiles,
	})
}

func GetProfileById(ctx *fiber.Ctx) error {
	collection := database.GetDBCollection("profiles")

	profileId := ctx.Params("profileId")

	objectID, err := primitive.ObjectIDFromHex(profileId)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid profile ID",
			"error":   err.Error(),
		})
	}

	var profile entity.Profile

	err = collection.FindOne(ctx.Context(), primitive.M{"_id": objectID}).Decode(&profile)
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Profile not found",
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Success get profile",
		"data":    profile,
	})
}

func DeleteProfile(ctx *fiber.Ctx) error {
	collection := database.GetDBCollection("profiles")

	profileId := ctx.Params("profileId")

	objectID, err := primitive.ObjectIDFromHex(profileId)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid profile ID",
			"error":   err.Error(),
		})
	}

	_, err = collection.DeleteOne(ctx.Context(), primitive.M{"_id": objectID})
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to delete profile",
			"error":   err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Success delete profile",
	})
}
