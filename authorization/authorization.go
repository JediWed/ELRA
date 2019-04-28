package authorization

import (
	"ELRA/globals"
	"ELRA/structs"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func CreateJWTToken(userid int, userrole int, configuration structs.Configuration) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["sub"] = userid
	claims["role"] = userrole
	claims["exp"] = time.Now().Add(time.Minute * 60).Unix()

	tokenString, err := token.SignedString([]byte(configuration.SigningKey))

	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func CheckAuthorization(endpoint func(http.ResponseWriter, *http.Request)) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		paramToken := strings.Split(r.Header.Get("Authorization"), "Bearer")
		if len(paramToken) == 2 {
			tokenString := strings.Trim(paramToken[1], " ")
			token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("There was an error")
				}
				return []byte(globals.Config.SigningKey), nil
			})

			if err != nil {
				w.WriteHeader(http.StatusUnauthorized)
			}

			if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
				log.Print("User with ID ", claims["sub"], " accessed ", r.RequestURI)
				endpoint(w, r)
			}
		} else {
			w.WriteHeader(http.StatusUnauthorized)
		}
	})
}

func ParseUserIDAndRole(request *http.Request) (int, int, error) {
	paramToken := strings.Split(request.Header.Get("Authorization"), "Bearer")
	if len(paramToken) == 2 {
		tokenString := strings.Trim(paramToken[1], " ")
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("There was an error")
			}
			return []byte(globals.Config.SigningKey), nil
		})

		if err != nil {
			return -1, -1, err
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			return int(claims["sub"].(float64)), int(claims["role"].(float64)), nil
		} else {
			return -1, -1, fmt.Errorf("Could not parse UserID")
		}
	}
	return -1, -1, fmt.Errorf("Could not parse Bearer Token")
}
