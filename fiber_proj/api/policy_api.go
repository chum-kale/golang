package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/your-package-name/controller" // Import your controller package
)

func SetupPolicyAPI(app *fiber.App, policyController *controller.PolicyController) {
	policyAPI := app.Group("/api/policy")

	policyAPI.Get("/list", policyController.ListPolicies)
	policyAPI.Get("/:policyName", policyController.GetPolicy)
}
