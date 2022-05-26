package domain

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/4kord/go-react-auth/internal/errs"
	"github.com/4kord/go-react-auth/internal/logger"
	"github.com/golang-jwt/jwt"
)

type User struct {
    Id int
    Username string
    Password string
    Role string
    jwt.StandardClaims
}

func (d User) GenerateToken() (string, *errs.Error) {
    atExp, err := strconv.Atoi(os.Getenv("AT_EXP"))
    if err != nil {
        logger.ErrorLog.Println(err.Error())
        return "", errs.ServerError("Error generating jwt")
    }

    claims := jwt.StandardClaims{
        Subject: strconv.Itoa(d.Id),
        ExpiresAt: time.Now().Add(time.Minute * time.Duration(atExp)).Unix(),
    }

    d.StandardClaims = claims

    aT := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

    aTString, err := aT.SignedString([]byte(os.Getenv("SECRET_KEY")))
    if err != nil {
        logger.ErrorLog.Println(err.Error())
        return "", errs.ServerError("Error generating jwt")
    }

    return aTString, nil
}

func (d User) ValidateToken(aT string) (bool, *errs.Error) {
    token, err := jwt.ParseWithClaims(aT, &d.StandardClaims, func(t *jwt.Token) (interface{}, error) {
        if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
            return nil, fmt.Errorf("Invalid signing method")
        }
        return []byte(os.Getenv("SECRET_KEY")), nil
    })
    if err != nil {
        logger.ErrorLog.Println(err.Error())
        return false, errs.ServerError("Error validating jwt")
    }

    return token.Valid, nil

}
