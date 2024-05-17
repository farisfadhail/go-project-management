package handlers

import (
	"github.com/gofiber/fiber/v2"
	"go-project-management/database"
	"go-project-management/models/entity"
	"go.mongodb.org/mongo-driver/bson"
)

func GetAllTasks(ctx *fiber.Ctx) error {
	collection := database.GetDBCollection("tasks")

	var tasks []entity.Task

	cursor, err := collection.Find(ctx.Context(), bson.M{})
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to get all tasks",
			"error":   err.Error(),
		})
	}

	if cursor.RemainingBatchLength() == 0 {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "No users found",
		})
	}

	err = cursor.All(ctx.Context(), &tasks)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to get all tasks",
			"error":   err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Success get all users",
		"data":    tasks,
	})
}

func GetTaskById(ctx *fiber.Ctx) error {
	taskId := ctx.Params("taskId")

	collection := database.GetDBCollection("tasks")

	var task entity.Task

	err := collection.FindOne(ctx.Context(), bson.M{"_id": taskId}).Decode(&task)
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Task not found",
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Success get task",
		"data":    task,
	})
}
