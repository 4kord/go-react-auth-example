package app

import (
	"database/sql"

	usersrepo "github.com/4kord/go-react-auth/internal/core/repositories/users"
	usersservice "github.com/4kord/go-react-auth/internal/core/services/users"
	"github.com/4kord/go-react-auth/internal/http/controllers"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/csrf"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

type Config struct {
    Db *sql.DB
    Port string
}

func Run(c Config) {
	userController := controllers.UserController{
		Service: usersservice.New(usersrepo.New(c.Db)),
	}

    app := fiber.New(fiber.Config{
        ErrorHandler: func(c *fiber.Ctx, err error) error {
            code := fiber.StatusInternalServerError
            if e, ok := err.(*fiber.Error); ok {
                code = e.Code
            }

            c.Status(code)
            err = c.JSON(fiber.Map{
                "status": "error",
                "code": code,
                "message": err.Error(),
            })
            if err != nil {
                return c.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
            }

            return nil
        },
    })

    app.Use(recover.New())
    app.Use(logger.New())
    app.Use(cors.New())
    app.Use(csrf.New())

    app.Post("/api/v1/auth/register", userController.Register)

    app.Listen(c.Port)
}
