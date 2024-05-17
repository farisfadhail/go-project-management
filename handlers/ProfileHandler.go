package handlers

import "github.com/gofiber/fiber/v2"

func CreateProfile(ctx *fiber.Ctx) error {
	return ctx.SendString("Create Profile")
}

func GetProfileById(ctx *fiber.Ctx) error {
	return ctx.SendString("Get Profile by ID")
}

func UpdateProfile(ctx *fiber.Ctx) error {
	return ctx.SendString("Update Profile")
}
