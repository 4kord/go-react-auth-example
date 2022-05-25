package main

import (
	"database/sql"
	"net/http"
	"os"

	"github.com/4kord/go-react-auth/internal/controller"
	"github.com/4kord/go-react-auth/internal/domain"
	"github.com/4kord/go-react-auth/internal/logger"
	"github.com/4kord/go-react-auth/internal/service"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
    _ "github.com/lib/pq"
)

func main() {
	err := godotenv.Load("/Users/4kord/Documents/Projects/go-react-auth/api/config/.env")
	if err != nil {
		logger.ErrorLog.Fatal(err)
	}

	dbConn, err := sql.Open(os.Getenv("DB_DRIVER"), os.Getenv("DB_DSN"))
	if err != nil {
		logger.ErrorLog.Fatal(err)
	}

	err = dbConn.Ping()
	if err != nil {
		logger.ErrorLog.Fatal(err)
	}

	userController := controller.UserController{
		Service: service.NewDefaultUserService(domain.NewUserRepositoryAdapter(dbConn)),
	}

	router := mux.NewRouter()

	router.HandleFunc("/api/v1/register", userController.Register)

	logger.InfoLog.Println("Starting server...")
	http.ListenAndServe(os.Getenv("PORT"), router)
}
