package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"go-project-management/database"
	"go-project-management/models/entity"
	"go-project-management/models/request"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

// Belum ditambahkan validasi admin project dengan session

func CreateProject(ctx *fiber.Ctx) error {
	registerRequest := new(request.ProjectRequest)
	err := ctx.BodyParser(registerRequest)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request",
			"error":   err.Error(),
		})
	}

	newProject := entity.Project{
		ID:          primitive.NewObjectID(),
		Title:       registerRequest.Title,
		Description: registerRequest.Description,
		Status:      registerRequest.Status,
		AdminUserId: registerRequest.AdminUserId, // data userId dari session
		StartDate:   registerRequest.StartDate,
		EndDate:     registerRequest.EndDate,
		Users:       registerRequest.Users,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	collection := database.GetDBCollection("projects")

	result, err := collection.InsertOne(ctx.Context(), newProject)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to create project",
			"error":   err.Error(),
		})
	}

	log.Info("Project created: ", result.InsertedID)

	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Success create project",
		"data":    newProject,
	})
}

func GetProjectById(ctx *fiber.Ctx) error {
	collection := database.GetDBCollection("projects")

	projectId := ctx.Params("projectId")

	objectId, err := primitive.ObjectIDFromHex(projectId)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid user ID",
			"error":   err.Error(),
		})
	}

	var project entity.Project

	err = collection.FindOne(ctx.Context(), bson.M{"_id": objectId}).Decode(&project)
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "User not found",
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Success get user by ID",
		"data":    project,
	})
}

func UpdateProject(ctx *fiber.Ctx) error {
	updateRequest := new(request.ProjectUpdateRequest)
	err := ctx.BodyParser(updateRequest)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request",
			"error":   err.Error(),
		})
	}

	collection := database.GetDBCollection("projects")

	projectId := ctx.Params("projectId")

	objectId, err := primitive.ObjectIDFromHex(projectId)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid project ID",
			"error":   err.Error(),
		})
	}

	var project entity.Project

	err = collection.FindOne(ctx.Context(), bson.M{"_id": objectId}).Decode(&project)
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Project not found",
		})
	}

	project.Title = updateRequest.Title
	project.Description = updateRequest.Description
	project.Status = updateRequest.Status
	project.EndDate = updateRequest.EndDate
	project.Users = updateRequest.Users
	project.UpdatedAt = time.Now()

	_, err = collection.UpdateOne(ctx.Context(), bson.M{"_id": objectId}, bson.M{"$set": project})
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to update project",
			"error":   err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Success update project",
		"data":    project,
	})
}

func DeleteProject(ctx *fiber.Ctx) error {
	collection := database.GetDBCollection("projects")

	projectId := ctx.Params("projectId")

	objectId, err := primitive.ObjectIDFromHex(projectId)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid project ID",
			"error":   err.Error(),
		})
	}

	deleteOne, err := collection.DeleteOne(ctx.Context(), bson.M{"_id": objectId})
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to delete project",
			"error":   err.Error(),
		})
	}

	if deleteOne.DeletedCount == 0 {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Project not found",
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": `Success delete project with ID ` + projectId,
	})
}
