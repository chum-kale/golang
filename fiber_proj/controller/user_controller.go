package controller

import (
	"github.com/gofiber/fiber/v2"
)

type UserController struct {
}

func NewUserController() *UserController {
	return &UserController{}
}

func (c *UserController) GetUser(ctx *fiber.Ctx) error {
	// Your implementation from the original main.go
	return nil
}

func (c *UserController) ListUsers(ctx *fiber.Ctx) error {
	// Your implementation from the original list_users.go
	return nil
}
