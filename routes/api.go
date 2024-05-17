package routes

import (
	"github.com/gofiber/fiber/v2"
	userHandlers "go-project-management/handlers"
	adminHandlers "go-project-management/handlers/admin"
	authHandlers "go-project-management/handlers/auth"
)

func RouteInit(app *fiber.App) {
	api := app.Group("/api")

	api.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.SendString("Welcome to Go Quiz API")
	})

	// Auth Routes
	api.Post("/login", authHandlers.Login).Name("login")
	api.Post("/register", authHandlers.Register).Name("register")

	admin := api.Group("/admin")

	// Admin User Routes
	adminUser := admin.Group("/user")
	adminUser.Get("/", adminHandlers.GetAllUsers).Name("admin.user.index")
	adminUser.Get("/:userId", adminHandlers.GetUserById).Name("admin.user.show")
	adminUser.Delete("/:userId", adminHandlers.DeleteUser).Name("admin.user.delete")

	// User Routes
	user := api.Group("/user")
	user.Get("/:userId", userHandlers.GetUserById).Name("user.show")
	user.Put("/:userId", userHandlers.UpdateUser).Name("user.update")

	// Admin Project Routes
	adminProject := admin.Group("/project")
	adminProject.Get("/", adminHandlers.GetAllProjects).Name("admin.project.index")
	adminProject.Get("/:projectId", adminHandlers.GetProjectById).Name("admin.project.show")

	// Project Routes
	project := api.Group("/project")
	project.Post("/", userHandlers.CreateProject).Name("project.create")
	project.Get("/:projectId", userHandlers.GetProjectById).Name("project.show")
	project.Put("/:projectId", userHandlers.UpdateProject).Name("project.update")
	project.Delete("/:projectId", userHandlers.DeleteProject).Name("project.delete")

	// Task Routes
	task := api.Group("/task")
	task.Post("/", userHandlers.CreateTask).Name("task.create")
	task.Get("/:taskId", userHandlers.GetTaskById).Name("task.show")
	task.Put("/:taskId", userHandlers.UpdateTask).Name("task.update")
	task.Delete("/:taskId", userHandlers.DeleteTask).Name("task.delete")

	// Meeting Routes
	meeting := api.Group("/meeting")
	meeting.Post("/", userHandlers.CreateMeeting).Name("meeting.create")
	meeting.Get("/:meetingId", userHandlers.GetMeetingById).Name("meeting.show")
	meeting.Put("/:meetingId", userHandlers.UpdateMeeting).Name("meeting.update")
	meeting.Delete("/:meetingId", userHandlers.DeleteMeeting).Name("meeting.delete")
}
