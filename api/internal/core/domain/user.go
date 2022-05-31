package domain

import (
	"os"
	"strconv"
	"time"

	"github.com/4kord/go-react-auth/internal/errs"
	"github.com/4kord/go-react-auth/internal/logger"
	"github.com/golang-jwt/jwt"
)

type User struct {
	Id       int
	Username string
	Password string
	Role     string
}

func (d *User) GenerateToken() (string, *errs.Error) {
	atExp, err := strconv.Atoi(os.Getenv("AT_EXP"))
	if err != nil {
		logger.ErrorLog.Println(err.Error())
		return "", errs.ServerError("Error generating jwt")
	}

	claims := jwt.StandardClaims{
		Subject:   strconv.Itoa(d.Id),
		ExpiresAt: time.Now().Add(time.Minute * time.Duration(atExp)).Unix(),
	}

	aT := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	aTString, err := aT.SignedString([]byte(os.Getenv("SECRET_KEY")))
	if err != nil {
		logger.ErrorLog.Println(err.Error())
		return "", errs.ServerError("Error generating jwt")
	}

	return aTString, nil
}

func (d *User) ValidateRole(role string) bool {
	return role == d.Role
}
