package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/your-package-name/service" // Import your IAM service package
)

type PolicyController struct {
	iamService *service.IAMService // Initialize IAM service in your controller
}

func NewPolicyController(iamService *service.IAMService) *PolicyController {
	return &PolicyController{
		iamService: iamService,
	}
}

func (c *PolicyController) ListPolicies(ctx *fiber.Ctx) error {
	policies, err := c.iamService.ListPolicies()

	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to fetch policies",
		})
	}

	return ctx.JSON(policies)
}

func (c *PolicyController) GetPolicy(ctx *fiber.Ctx) error {
	policyName := ctx.Params("policyName")
	policy, err := c.iamService.GetPolicy(policyName)

	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to fetch policy",
		})
	}

	return ctx.JSON(policy)
}
