package handlers

import (
	"github.com/gofiber/fiber/v2"
	"go-project-management/database"
	"go-project-management/models/entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetAllUsers(ctx *fiber.Ctx) error {
	collection := database.GetDBCollection("users")

	var users []entity.User

	cursor, err := collection.Find(ctx.Context(), bson.M{})
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to get all users",
			"error":   err.Error(),
		})
	}

	if cursor.RemainingBatchLength() == 0 {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "No users found",
		})
	}

	err = cursor.All(ctx.Context(), &users)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to get all users",
			"error":   err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Success get all users",
		"data":    users,
	})
}

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

func DeleteUser(ctx *fiber.Ctx) error {
	collection := database.GetDBCollection("users")

	userId := ctx.Params("userId")

	deleteOne, err := collection.DeleteOne(ctx.Context(), bson.M{"_id": userId})
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to delete user",
			"error":   err.Error(),
		})
	}

	if deleteOne.DeletedCount == 0 {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "User not found",
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": `Success delete user with ID ` + userId,
	})
}
