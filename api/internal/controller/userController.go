package controller

import (
	"encoding/json"
	"net/http"

	"github.com/4kord/go-react-auth/internal/dto"
	"github.com/4kord/go-react-auth/internal/logger"
	"github.com/4kord/go-react-auth/internal/service"
)

type UserController struct{
    Service service.UserService
}

func (c UserController) Register(w http.ResponseWriter, r *http.Request) {
    var request dto.UserRequest

    err := json.NewDecoder(r.Body).Decode(&request)
    if err != nil {
        logger.ErrorLog.Println(err.Error())
		writeResponse(w, http.StatusBadRequest, err.Error())
        return
    }

    e := c.Service.Register(request)
    if e != nil {
        logger.ErrorLog.Println(e.Message)
		writeResponse(w, e.Code, e)
        return
    }

    writeResponse(w, http.StatusCreated, nil)
}
