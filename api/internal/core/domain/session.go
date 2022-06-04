package domain

import (
	"os"
	"strconv"
	"time"

	"github.com/4kord/go-react-auth/internal/errs"
	"github.com/4kord/go-react-auth/internal/logger"
	"github.com/golang-jwt/jwt"
)

type Session struct {
	Id           int
	UserId       int
	RefreshToken string
	Ip           string
	ExpiresAt    time.Time
	CreatedAt    time.Time
}

func (d *Session) ValidateExpiry() bool {
	return !time.Now().UTC().After(d.ExpiresAt)
}

func (d *Session) ValidateIp(requestIp string) bool {
	return d.Ip == requestIp

}

func (d *Session) GenerateToken() (string, *errs.Error) {
	atExp, err := strconv.Atoi(os.Getenv("AT_EXP"))
	if err != nil {
		logger.ErrorLog.Println(err.Error())
		return "", errs.ServerError("Error generating jwt")
	}

	claims := jwt.StandardClaims{
		Subject:   strconv.Itoa(d.UserId),
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
