package helper

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/kodinggo/golang-bootcamp-batch2/rest-api-service/internal/config"
	"github.com/kodinggo/golang-bootcamp-batch2/rest-api-service/internal/model"
)

func GenerateToken(userID int64) (strToken string, err error) {
	expiredAt := time.Now().UTC().Add(config.JWTExp())
	strToken, err = jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp":     expiredAt.Unix(),
		"user_id": userID,
	}).SignedString([]byte(config.JWTSigningKey()))
	return
}

func DecodeToken(token string, claim *model.CustomClaims) (err error) {
	jwt.ParseWithClaims(token, claim, func(t *jwt.Token) (interface{}, error) {
		return []byte(config.JWTSigningKey()), nil
	})
	return
}
