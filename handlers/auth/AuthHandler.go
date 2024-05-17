package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"go-project-management/database"
	"go-project-management/models/entity"
	"go-project-management/models/request"
	"go-project-management/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

func Register(ctx *fiber.Ctx) error {
	registerRequest := new(request.RegisterRequest)
	err := ctx.BodyParser(registerRequest)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request",
			"error":   err.Error(),
		})
	}

	if registerRequest.Password != registerRequest.ConfirmPassword {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Password and Confirm Password must be the same",
		})
	}

	password, err := utils.HashingPassword(registerRequest.Password)

	newUser := entity.User{
		ID:        primitive.NewObjectID(),
		Username:  registerRequest.Username,
		Email:     registerRequest.Email,
		Password:  password,
		Role:      "user",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	collection := database.GetDBCollection("users")

	err = collection.FindOne(ctx.Context(), bson.M{"email": newUser.Email}).Decode(&newUser)
	if err == nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Email already registered",
		})
	}

	err = collection.FindOne(ctx.Context(), bson.M{"username": newUser.Username}).Decode(&newUser)
	if err == nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Username already registered",
		})
	}

	result, err := collection.InsertOne(ctx.Context(), newUser)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to register user",
		})
	}

	log.Info("User created: ", result.InsertedID)

	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "User registered successfully",
		"data":    result,
	})
}

func Login(ctx *fiber.Ctx) error {
	return nil
}
