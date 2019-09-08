package middlewares

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
)

func ValidateUser(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authorization := r.Header.Get("Authorization")

		if authorization != "" {
			bearerToken := strings.Split(authorization, " ")

			if len(bearerToken) == 2 {
				token, err := jwt.Parse(bearerToken[1], func(token *jwt.Token) (interface{}, error) {
					if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
						return nil, fmt.Errorf("There was an error")
					}

					return []byte("secret"), nil
				})

				if err != nil {
					w.WriteHeader(http.StatusInternalServerError)
					w.Write([]byte(`{"message": "` + err.Error() + `"}`))

					return
				}

				if token.Valid {
					next(w, r)
				}
			} else {
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte(`{"message": "Unauthorized"}`))

				return
			}
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(`{"message": "Unauthorized"}`))

			return
		}
	})
}
