package routes

import (
	"database/sql"

	usersrepo "github.com/4kord/go-react-auth/internal/core/repositories/users"
	usersservice "github.com/4kord/go-react-auth/internal/core/services/users"
	"github.com/4kord/go-react-auth/internal/http/controllers"
	"github.com/gorilla/mux"
)

func Routes(db *sql.DB) *mux.Router {
	userController := controllers.UserController{
		Service: usersservice.New(usersrepo.New(db)),
	}

	router := mux.NewRouter()

	router.HandleFunc("/api/v1/register", userController.Register)

    return router
}
