package utils

import (
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/Minhvn98/ecommerce-fashion/config"
	"github.com/Minhvn98/ecommerce-fashion/models"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		log.Fatal(err)
	}
	return string(bytes)
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func SetCookieHandler(w http.ResponseWriter, r *http.Request, name string, value string) {
	cookie := &http.Cookie{
		Name:     name,
		Value:    value,
		Expires:  time.Now().Add(time.Hour * 24),
		HttpOnly: true,
	}
	http.SetCookie(w, cookie)
}

func ReadCookieHandler(w http.ResponseWriter, r *http.Request, name string) string {
	cookie, err := r.Cookie(name)
	if nil != err {
		log.Fatal(err)
	}
	return cookie.Value
}

func CreateToken(user models.User) (string, error) {
	claims := models.Claims{
		user.Name,
		user.Role,
		jwt.StandardClaims{
			Issuer:    strconv.Itoa(int(user.ID)),
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(), //1 day
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenStr, err := token.SignedString([]byte(config.Config.AccessSecret))
	if nil != err {
		log.Fatal(err)
	}
	return tokenStr, err
}

func VerifyAndParserToken(tokenStr string) (bool, models.Claims) {
	token, _ := jwt.ParseWithClaims(tokenStr, &models.Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.Config.AccessSecret), nil
	})

	if claims, ok := token.Claims.(*models.Claims); ok && token.Valid {
		return token.Valid, *claims
	} else {
		return token.Valid, models.Claims{}
	}
}
