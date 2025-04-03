package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/your-package-name/controller" // Import your controller package
)

func SetupUserAPI(app *fiber.App, userController *controller.UserController) {
	userAPI := app.Group("/api/user")

	userAPI.Get("/:username", userController.GetUser)
	userAPI.Get("/list", userController.ListUsers)
}
