package app

import (
	"context"
	"mobile-rest-api/config"
	"mobile-rest-api/models"
	"mobile-rest-api/utils"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
)

//NotAuthURls array
var NotAuthURls = []string{
	"/",
}

//JWTAutentication middleware
var JWTAutentication = func(nextHandler http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			var requestPath = r.URL.Path

			for _, value := range NotAuthURls {
				if value == requestPath {
					nextHandler.ServeHTTP(w, r)
					return
				}
			}
			var linkPublic = string([]rune(requestPath)[0:8])
			if linkPublic == "/public/" {
				nextHandler.ServeHTTP(w, r)
				return
			}

			var response = make(map[string]interface{})

			var tokenHeader = r.Header.Get("Authorization")

			if tokenHeader == "" {
				response = utils.Message(730, "Отсутствует токен авторизации")
				w.WriteHeader(http.StatusForbidden)
				utils.Respond(w, response)
				return
			}

			var splitted = strings.Split(tokenHeader, " ")
			if len(splitted) != 2 {
				response = utils.Message(731, "Неверный / неправильно сформированный токен авторизации")
				w.WriteHeader(http.StatusForbidden)
				utils.Respond(w, response)
				return
			}

			var tokenPart = splitted[1]
			var tk = &models.Token{}

			var token, err = jwt.ParseWithClaims(tokenPart, tk, func(token *jwt.Token) (interface{}, error) {

				config.Init()
				var cfg = config.Get()
				return []byte(cfg.SecretKey), nil
			})
			if err != nil {
				response = utils.Message(732, "Неверно сформированный токен аутентификации")
				w.WriteHeader(http.StatusForbidden)
				utils.Respond(w, response)
				return
			}
			if !token.Valid {
				response = utils.Message(732, "Токен недействителен.")
				w.WriteHeader(http.StatusForbidden)
				utils.Respond(w, response)
				return
			}
			type ctxUser struct {
				ID      uint64
				phoneNo string
				smsCode string
			}
			var ctU ctxUser
			ctU.ID = tk.UserID
			ctU.phoneNo = tk.PhoneNo
			ctU.smsCode = tk.SmsCode
			//log.Println(" Logined user ID: ", ctU)

			var ctx = context.WithValue(r.Context(), "user", ctU)

			r = r.WithContext(ctx)
			nextHandler.ServeHTTP(w, r)
			return

		})

}
