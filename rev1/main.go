package main

import (
	"/work/golang/rev1/docs"

	"github.com/gofiber/fiber/v2"
	"github.com/swaggo/fiber-swagger"
)

func main() {
	app := fiber.New()

	// Register the Swagger documentation route
	app.Get("/swagger/*", swagger.New(swagger.Config{ // This route will serve the generated Swagger documentation
		URL: "/docs/doc.json", // The URL to the generated Swagger JSON file
	}))

	// Initialize the API routes
	userController := controllers.NewUserController()
	roleController := controllers.NewRoleController()
	policyController := controllers.NewPolicyController()

	app.Get("/api/user/:username", userController.GetUser)
	app.Get("/api/users", userController.ListUsers)
	app.Get("/api/policies", policyController.GetPolicies)
	app.Get("/api/roles/:maxRoles", roleController.ListRoles)
	app.Post("/api/user", userController.CreateUser)
	app.Get("/api/user/:username/exists", userController.CheckUserExists)
	app.Get("/api/group/:group/exists", userController.CheckGroupExists)

	// Serve Swagger JSON on a specific route
	docs.SwaggerInfo.BasePath = "/api" // Set your base path
	app.Get("/docs/*", func(c *fiber.Ctx) {
		c.JSON(docs.SwaggerInfo) // Serve the Swagger JSON info
	})

	app.Listen(":3000")
}
