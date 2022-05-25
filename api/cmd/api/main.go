package main

import (
	"database/sql"
	"net/http"
	"os"

	"github.com/4kord/go-react-auth/internal/controllers"
	"github.com/4kord/go-react-auth/internal/core/repositories/usersrepo"
	"github.com/4kord/go-react-auth/internal/core/services/users"
	"github.com/4kord/go-react-auth/internal/logger"
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

	userController := controllers.UserController{
		Service: users.New(usersrepo.New(dbConn)),
	}

	router := mux.NewRouter()

	router.HandleFunc("/api/v1/register", userController.Register)

	logger.InfoLog.Println("Starting server...")
	http.ListenAndServe(os.Getenv("PORT"), router)
}
