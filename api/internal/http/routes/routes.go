package routes

import (
	"database/sql"

	"github.com/4kord/go-react-auth/internal/core/repositories/sessions"
	"github.com/4kord/go-react-auth/internal/core/repositories/users"
	authservice "github.com/4kord/go-react-auth/internal/core/services/auth"
	usersservice "github.com/4kord/go-react-auth/internal/core/services/users"
	"github.com/4kord/go-react-auth/internal/http/controllers"
	"github.com/4kord/go-react-auth/internal/http/middlewares/auth"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func Setup(app *fiber.App, db *sql.DB) *fiber.App {
	authController := controllers.AuthController{
		Service: authservice.New(users.New(db), sessions.New(db)),
	}

	userController := controllers.UserController{
		Service: usersservice.New(users.New(db)),
	}

	app.Use(recover.New())
	app.Use(logger.New())
	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))

	app.Use("/api/admin", auth.New(auth.Config{
		Repo: users.New(db),
		Role: "admin",
	}))

	app.Use("/api/user", auth.New(auth.Config{
		Repo: users.New(db),
		Role: "user",
	}))

	app.Post("/api/auth/login", authController.Login)
	app.Post("/api/auth/register", authController.Register)
	app.Post("/api/auth/logout", authController.Logout)
	app.Get("/api/auth/refresh", authController.Refresh)

	app.Get("/api/user/me", userController.Me)

	app.Get("/api/user/test", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "SUPER SECRET USER INFO",
		})
	})

	return app
}
