package main

import (
	"os"

	"github.com/4kord/go-react-auth/internal/http/db"
	"github.com/4kord/go-react-auth/internal/http/routes"
	"github.com/4kord/go-react-auth/internal/logger"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)


func main() {
	err := godotenv.Load("/Users/4kord/Documents/Projects/go-react-auth/api/config/.env")
	if err != nil {
		logger.ErrorLog.Fatal(err)
	}

    db, err := db.Setup(os.Getenv("DB_DRIVER"), os.Getenv("DB_DSN"))
    if err != nil {
		logger.ErrorLog.Fatal(err)
    }
    defer db.Close()

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

    logger.ErrorLog.Fatal(routes.Setup(app, db).Listen(os.Getenv("PORT")))
}
