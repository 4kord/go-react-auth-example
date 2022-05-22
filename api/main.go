package main

import (
	"os"

	"github.com/joho/godotenv"
)

type config struct {
    port string
    dbDriver string
    dsn string
}

func main() {
    godotenv.Load()
    config := config{
        port: os.Getenv("PORT"),
        dbDriver: os.Getenv("DB_DRIVER"),
        dsn: os.Getenv("DSN"),
    }
}
