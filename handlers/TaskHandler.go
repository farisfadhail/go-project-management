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

func CreateTask(ctx *fiber.Ctx) error {
	taskRequest := new(request.TaskRequest)
	err := ctx.BodyParser(taskRequest)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request",
			"error":   err.Error(),
		})
	}

	findOne := database.GetDBCollection("projects").FindOne(ctx.Context(), primitive.M{"_id": taskRequest.ProjectId})
	if findOne.Err() != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Project not found",
			"error":   findOne.Err().Error(),
		})
	}

	newTask := entity.Task{
		ID:              primitive.NewObjectID(),
		ProjectId:       taskRequest.ProjectId,
		TaskTitle:       taskRequest.TaskTitle,
		TaskDescription: taskRequest.TaskDescription,
		TaskStatus:      taskRequest.TaskStatus,
		Users:           taskRequest.Users,
		CreatedAt:       time.Now(),
		UpdatedAt:       time.Now(),
	}

	collection := database.GetDBCollection("tasks")

	result, err := collection.InsertOne(ctx.Context(), newTask)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to create task",
			"error":   err.Error(),
		})
	}

	log.Info("Task created: ", result.InsertedID)

	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Success create task",
		"data":    newTask,
	})
}

func GetTaskById(ctx *fiber.Ctx) error {
	collection := database.GetDBCollection("tasks")

	taskId := ctx.Params("taskId")

	objectId, err := primitive.ObjectIDFromHex(taskId)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid task id",
			"error":   err.Error(),
		})
	}

	var task entity.Task

	err = collection.FindOne(ctx.Context(), bson.M{"_id": objectId}).Decode(&task)
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

func UpdateTask(ctx *fiber.Ctx) error {
	updateRequest := new(request.TaskUpdateRequest)
	err := ctx.BodyParser(updateRequest)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request",
			"error":   err.Error(),
		})
	}

	collection := database.GetDBCollection("tasks")

	taskId := ctx.Params("taskId")

	objectID, err := primitive.ObjectIDFromHex(taskId)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid task ID",
			"error":   err.Error(),
		})
	}

	var task entity.Task

	err = collection.FindOne(ctx.Context(), bson.M{"_id": objectID}).Decode(&task)
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Task not found",
		})
	}

	task.TaskTitle = updateRequest.TaskTitle
	task.TaskDescription = updateRequest.TaskDescription
	task.TaskStatus = updateRequest.TaskStatus
	task.Users = updateRequest.Users
	task.UpdatedAt = time.Now()

	_, err = collection.UpdateOne(ctx.Context(), bson.M{"_id": objectID}, bson.M{"$set": task})
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to update task",
			"error":   err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Success update task",
		"data":    task,
	})
}

func DeleteTask(ctx *fiber.Ctx) error {
	collection := database.GetDBCollection("tasks")

	taskId := ctx.Params("taskId")

	objectId, err := primitive.ObjectIDFromHex(taskId)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid task ID",
			"error":   err.Error(),
		})
	}

	deleteOne, err := collection.DeleteOne(ctx.Context(), bson.M{"_id": objectId})
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to delete task",
			"error":   err.Error(),
		})
	}

	if deleteOne.DeletedCount == 0 {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Task not found",
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Success delete task",
	})
}
