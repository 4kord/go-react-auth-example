package controllers

import (
	"net/http"

	"github.com/4kord/go-react-auth/internal/core/services/auth"
	"github.com/4kord/go-react-auth/internal/dto"
	"github.com/4kord/go-react-auth/internal/logger"
	"github.com/gofiber/fiber/v2"
)

type UserController struct {
	Service auth.Service
}

func (ctrl UserController) Register(c *fiber.Ctx) error {
	var request dto.UserRequest

	err := c.BodyParser(&request)
	if err != nil {
		logger.ErrorLog.Println(err.Error())
		return fiber.NewError(http.StatusBadRequest, err.Error())
	}

	e := ctrl.Service.Register(request)
	if e != nil {
		logger.ErrorLog.Println(e.Message)
		return fiber.NewError(e.Code, e.Message)
	}

	c.Status(http.StatusCreated)
	return c.JSON(fiber.Map{
		"message": "Account created",
	})
}

func (ctrl UserController) Login(c *fiber.Ctx) error {
	var request dto.UserRequest

	err := c.BodyParser(&request)
	if err != nil {
		logger.ErrorLog.Println(err.Error())
		return fiber.NewError(http.StatusBadRequest, err.Error())
	}

    request.Ip = c.IP()

	response, e := ctrl.Service.Login(request)
	if e != nil {
		logger.ErrorLog.Println(e.Message)
		return fiber.NewError(e.Code, e.Message)
	}

	c.Status(http.StatusOK)
	return c.JSON(response)
}
