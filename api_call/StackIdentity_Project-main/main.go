package main

import (
	"StackIdentity_Project/controllers"
	_ "StackIdentity_Project/docs" // Import your auto-generated docs

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
)

func main() {
	app := fiber.New()

	userController := controllers.NewUserController()
	roleController := controllers.NewRoleController()
	policyController := controllers.NewPolicyController()

	// Register your API endpoints
	app.Get("/api/user/:username", userController.GetUser)
	app.Get("/api/users", userController.ListUsers)
	app.Get("/api/policies", policyController.GetPolicies)
	app.Get("/api/roles/:maxRoles", roleController.ListRoles)
	app.Post("/api/user", userController.CreateUser)

	// Add a route to check if a user exists by username
	app.Get("/api/user/:username/exists", userController.CheckUserExists)

	// Serve Swagger UI and JSON documentation
	app.Get("/swagger/*", swagger.HandlerDefault) // default
	app.Get("/swagger/*", swagger.New(swagger.Config{
		URL:          "/swagger/doc.json", // Use the relative path to your Swagger JSON documentation
		DeepLinking:  false,
		DocExpansion: "none",
		OAuth: &swagger.OAuthConfig{
			AppName:  "OAuth Provider",
			ClientId: "21bb4edc-05a7-4afc-86f1-2e151e4ba6e2",
		},
		OAuth2RedirectUrl: "http://localhost:3000/swagger/oauth2-redirect.html",
	}))

	// Start your Fiber application
	app.Listen(":3000")
}
