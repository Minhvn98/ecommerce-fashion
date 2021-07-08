package utils

import (
	"log"
	"net/http"
	"strconv"
	"time"

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

func CreateToken(idUser int) (string, error) {
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer:    strconv.Itoa(int(idUser)),
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(), //1 day
	})

	token, err := claims.SignedString([]byte("learnmorelearnmore"))
	if nil != err {
		log.Fatal(err)
	}
	return token, err
}
