package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/your-project/api"        // Import your API package
	"github.com/your-project/controller" // Import your controller package
)

func main() {
	// Create a Fiber app
	app := fiber.New()

	// Initialize your controller instances (assuming you have these defined)
	userController := controller.NewUserController()
	roleController := controller.NewRoleController()
	policyController := controller.NewPolicyController()

	// Configure routes using API setup functions
	api.SetupUserAPI(app, userController)
	api.SetupRoleAPI(app, roleController)
	api.SetupPolicyAPI(app, policyController)

	// Start the Fiber app on a specified port
	app.Listen(":8080") // You can change the port as needed
}
