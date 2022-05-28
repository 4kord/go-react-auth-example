package auth

import (
	"fmt"
	"os"

	"github.com/4kord/go-react-auth/internal/errs"
	"github.com/4kord/go-react-auth/internal/logger"
	"github.com/golang-jwt/jwt"
)

func validateToken(aT string) (*jwt.StandardClaims, *errs.Error) {
    var claims jwt.StandardClaims
    token, err := jwt.ParseWithClaims(aT, &claims, func(t *jwt.Token) (interface{}, error) {
        if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
            return nil, fmt.Errorf("Invalid signing method")
        }
        return []byte(os.Getenv("SECRET_KEY")), nil
    })
    if err != nil || !token.Valid {
        logger.ErrorLog.Println(err.Error())
        return nil, errs.UnAuthorizedError("Token isn't valid")
    }

    return &claims, nil
}
