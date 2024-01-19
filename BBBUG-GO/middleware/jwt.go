package middleware

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/msinwolc/config"
	"github.com/msinwolc/models"
)

var jwtKey = []byte("msinwolc")

func CreateToken(userName, plat string, userId int, expireDuration time.Duration) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"UserName": userName,
		"UserId":   userId,
		"UserPlat": plat,
		"iat":      time.Now().Unix(),
		"exp":      time.Now().Add(expireDuration).Unix(),
	})
	return token.SignedString(jwtKey)
}

type Claims struct {
	UserName string
	UserId   int
	UserPlat string
	jwt.StandardClaims
}

func parseToken(tokenstr string) (*jwt.Token, *Claims, error) {
	claims := new(Claims)
	token, err := jwt.ParseWithClaims(tokenstr, claims, func(t *jwt.Token) (interface{}, error) { return jwtKey, nil })
	return token, claims, err
}

func GetUserName(token string) (string, error) {
	_, claim, err := parseToken(token)
	if claim != nil && err == nil {
		return claim.UserName, nil
	}
	return "", err
}

func GetUserId(token string) (int, error) {
	_, claim, err := parseToken(token)
	if claim != nil && err == nil {
		return claim.UserId, nil
	}
	return 0, err
}

func CheckJwtAuth(str string) (bool, error) {
	token, claims, err := parseToken(str)
	if err != nil && token.Valid {
		rds_token, err := config.GetToken(claims.UserName)
		if err != nil {
			return false, err
		}
		if rds_token == token.Raw {
			status, err := models.GetUserStatusById(claims.UserId)
			if err != nil {
				return false, err
			}
			if status == 1 {
				return false, fmt.Errorf("账号已被锁定")
			}
			return true, nil
		}
	}
	return false, err
}
