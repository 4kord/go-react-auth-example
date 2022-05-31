package auth

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func New(config Config) fiber.Handler {
    return func(c *fiber.Ctx) error {
        authHeader := c.Get("Authorization")    
        if authHeader == "" {
            return fiber.NewError(http.StatusUnauthorized, "Authorization header can't be empty")
        }

        authHeaderSlice := strings.Split(authHeader, " ")
        if len(authHeaderSlice) != 2 || authHeaderSlice[0] != "Bearer" {
            return fiber.NewError(http.StatusUnauthorized, "Authorization header is in incorrect format")
        }

        claims, e := validateToken(authHeaderSlice[1])
        if e != nil {
            return fiber.NewError(e.Code, e.Message)
        }

        if config.Role == "" {
            return c.Next()
        }

        id, err := strconv.Atoi(claims.Subject)
        if err != nil {
            return fiber.NewError(http.StatusUnauthorized, "Error authorizing role")
        }

        user, e := config.Repo.GetUser(id)
        if e != nil {
            return fiber.NewError(e.Code, e.Message)
        }

        if !user.ValidateRole(config.Role) {
            return fiber.NewError(http.StatusForbidden, "You don't have permissions to see this page") 
        }

        return c.Next()
    }
}
