package main

import (
	"net/http"
	"os"

	"github.com/4kord/go-react-auth/internal/http/db"
	"github.com/4kord/go-react-auth/internal/http/routes"
	"github.com/4kord/go-react-auth/internal/logger"
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

    logger.ErrorLog.Fatalln(http.ListenAndServe(os.Getenv("PORT"), routes.Routes(db)))
}
