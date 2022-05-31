package controllers

import (
	"net/http"
	"time"

	"github.com/4kord/go-react-auth/internal/core/services/auth"
	"github.com/4kord/go-react-auth/internal/dto"
	"github.com/gofiber/fiber/v2"
)

type AuthController struct {
	Service auth.Service
}

func (ctrl AuthController) Register(c *fiber.Ctx) error {
	var request dto.UserRequest

	err := c.BodyParser(&request)
	if err != nil {
		return fiber.NewError(http.StatusBadRequest, err.Error())
	}

	e := ctrl.Service.Register(request)
	if e != nil {
		return fiber.NewError(e.Code, e.Message)
	}

	c.Status(http.StatusCreated)
	return c.JSON(fiber.Map{
		"message": "Account created",
	})
}

func (ctrl AuthController) Login(c *fiber.Ctx) error {
	var request dto.UserRequest

	err := c.BodyParser(&request)
	if err != nil {
		return fiber.NewError(http.StatusBadRequest, err.Error())
	}

	request.Ip = c.IP()

	response, e := ctrl.Service.Login(request)
	if e != nil {
		return fiber.NewError(e.Code, e.Message)
	}

	c.Cookie(&fiber.Cookie{
		Name:     "session",
		Value:    response.RefreshToken,
		Expires:  response.RefreshExpires,
		HTTPOnly: true,
	})

	c.Status(http.StatusOK)
	return c.JSON(response)
}

func (ctrl AuthController) Logout(c *fiber.Ctx) error {
	ctrl.Service.Logout(dto.SessionRequest{
		RefreshToken: c.Cookies("session"),
	})

	c.Cookie(&fiber.Cookie{
		Name:    "session",
		MaxAge:  -1,
		Expires: time.Now().Add(-100 * time.Hour),
	})

	return c.JSON(fiber.Map{
		"message": "successfully logout",
	})
}

func (ctrl AuthController) Refresh(c *fiber.Ctx) error {
	request := dto.SessionRequest{
		RefreshToken: c.Cookies("session"),
		Ip:           c.IP(),
	}

	response, err := ctrl.Service.Refresh(request)
	if err != nil {
		return fiber.NewError(err.Code, err.Message)
	}

	c.Cookie(&fiber.Cookie{
		Name:     "session",
		Value:    response.RefreshToken,
		Expires:  response.RefreshExpires,
		HTTPOnly: true,
	})

	return c.JSON(response)
}
