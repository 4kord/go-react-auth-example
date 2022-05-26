package routes

import (
	"database/sql"

	usersrepo "github.com/4kord/go-react-auth/internal/core/repositories/users"
	usersservice "github.com/4kord/go-react-auth/internal/core/services/users"
	"github.com/4kord/go-react-auth/internal/http/controllers"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func Setup(app *fiber.App, db *sql.DB) *fiber.App {
	userController := controllers.UserController{
		Service: usersservice.New(usersrepo.New(db)),
	}

    app.Use(recover.New())
    app.Use(logger.New())
    app.Use(cors.New())
    
    app.Post("/api/auth/login", userController.Login)
    app.Post("/api/auth/register", userController.Register)

    return app
}
