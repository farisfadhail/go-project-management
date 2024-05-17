package handlers

import (
	"github.com/gofiber/fiber/v2"
	"go-project-management/database"
	"go-project-management/models/entity"
	"go.mongodb.org/mongo-driver/bson"
)

func GetAllProjects(ctx *fiber.Ctx) error {
	collection := database.GetDBCollection("projects")

	var projects []entity.Project

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

	err = cursor.All(ctx.Context(), &projects)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to get all users",
			"error":   err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Success get all users",
		"data":    projects,
	})
}

func GetProjectById(ctx *fiber.Ctx) error {
	projectId := ctx.Params("projectId")

	collection := database.GetDBCollection("projects")

	var project entity.Project

	err := collection.FindOne(ctx.Context(), bson.M{"_id": projectId}).Decode(&project)
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Project not found",
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Success get project",
		"data":    project,
	})
}
