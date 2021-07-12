package middlewares

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/Minhvn98/ecommerce-fashion/utils"
	"github.com/gorilla/context"
)

func CommonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "*")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, PATCH, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, Access-Control-Request-Headers, Access-Control-Request-Method, Connection, Host, Origin, User-Agent, Referer, Cache-Control, X-header")
		next.ServeHTTP(w, r)
	})
}

func Authentication(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		c, err := r.Cookie("token")
		if err != nil {
			fmt.Println(err)
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(map[string]interface{}{"message": "Invalid token"})
		} else {
			check, claims := utils.VerifyAndParserToken(c.Value)
			if check {
				context.Set(r, "id", claims.Issuer)
				context.Set(r, "role", claims.Role)
				next.ServeHTTP(w, r)
			} else {
				w.WriteHeader(http.StatusUnauthorized)
				json.NewEncoder(w).Encode(map[string]interface{}{"message": "Invalid token haha", "claims": claims})
			}
		}
	})
}

func GetUserIdFromToken(w http.ResponseWriter, r *http.Request, roles ...string) int {
	role := context.Get(r, "role")
	userId := context.Get(r, "id")
	for _, value := range roles {
		if value == role {
			i, _ := strconv.ParseInt(userId.(string), 10, 64)
			return int(i)
		}
	}

	return 0
}
