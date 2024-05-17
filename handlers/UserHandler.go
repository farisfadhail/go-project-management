package handlers

import (
	"github.com/gofiber/fiber/v2"
	"go-project-management/database"
	"go-project-management/models/entity"
	"go-project-management/models/request"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

func GetUserById(ctx *fiber.Ctx) error {
	collection := database.GetDBCollection("users")

	userId := ctx.Params("userId")

	objectId, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid user ID",
			"error":   err.Error(),
		})
	}

	var user entity.User

	err = collection.FindOne(ctx.Context(), bson.M{"_id": objectId}).Decode(&user)
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "User not found",
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Success get user by ID",
		"data":    user,
	})
}

func UpdateUser(ctx *fiber.Ctx) error {
	userUpdateRequest := new(request.UserUpdateRequest)
	err := ctx.BodyParser(userUpdateRequest)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request",
			"error":   err.Error(),
		})
	}

	collection := database.GetDBCollection("users")

	userId := ctx.Params("userId")

	objectId, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid user ID",
			"error":   err.Error(),
		})
	}

	var user entity.User

	err = collection.FindOne(ctx.Context(), bson.M{"_id": objectId}).Decode(&user)
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "User not found",
		})
	}

	user.Username = userUpdateRequest.Username
	user.UpdatedAt = time.Now()

	_, err = collection.UpdateOne(ctx.Context(), bson.M{"_id": objectId}, bson.M{"$set": user})
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to update user",
			"error":   err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Success update user",
		"data":    user,
	})
}

func UpdateEmailUser(ctx *fiber.Ctx) error {
	return nil
}
