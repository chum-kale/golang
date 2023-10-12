package controller

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/your-package-name/service" // Import your IAM service package
)

type RoleController struct {
	iamService *service.IAMService // Initialize IAM service in your controller
}

func NewRoleController(iamService *service.IAMService) *RoleController {
	return &RoleController{
		iamService: iamService,
	}
}

func (c *RoleController) ListRoles(ctx *fiber.Ctx) error {
	maxRolesParam := ctx.Params("maxRoles")
	maxRoles, err := strconv.Atoi(maxRolesParam)

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid maxRoles parameter",
		})
	}

	roles, err := c.iamService.ListRoles(maxRoles)

	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to fetch roles",
		})
	}

	return ctx.JSON(roles)
}
