package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/your-package-name/controller" // Import your controller package
)

func SetupRoleAPI(app *fiber.App, roleController *controller.RoleController) {
	roleAPI := app.Group("/api/role")

	roleAPI.Get("/list/:maxRoles", roleController.ListRoles)
}
