package controllers

import (
	"github.com/4kord/go-react-auth/internal/core/services/users"
	"github.com/gofiber/fiber/v2"
)

type UserController struct {
	Service users.Service
}

func (ctrl UserController) Me(c *fiber.Ctx) error {
	if userId, ok := c.Locals("userId").(int); ok {
		response, err := ctrl.Service.GetUser(userId)
		if err != nil {
			return fiber.NewError(err.Code, err.Message)
		}
		return c.JSON(response)
	}

	return fiber.NewError(500, "Unexpected error")
}
