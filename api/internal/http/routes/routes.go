package routes

import (
	"database/sql"

	"github.com/4kord/go-react-auth/internal/core/repositories/sessions"
    "github.com/4kord/go-react-auth/internal/core/repositories/users"
	authservice "github.com/4kord/go-react-auth/internal/core/services/auth"
	"github.com/4kord/go-react-auth/internal/http/controllers"
	"github.com/4kord/go-react-auth/internal/http/middlewares/auth"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func Setup(app *fiber.App, db *sql.DB) *fiber.App {
	userController := controllers.UserController{
		Service: authservice.New(users.New(db), sessions.New(db)),
	}

	app.Use(recover.New())
	app.Use(logger.New())
	app.Use(cors.New())

	app.Use("/api/user", auth.New(auth.Config{
		Repo: users.New(db),
		Role: "admin",
	}))

	app.Post("/api/auth/login", userController.Login)
	app.Post("/api/auth/register", userController.Register)

	app.Post("/api/user/test", func(c *fiber.Ctx) error {
		return c.SendString("test protected page")
	})

	return app
}
