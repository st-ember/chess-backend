package middleware

import (
	"errors"
	"net/http"

	log "github.com/sirupsen/logrus"
	"github.com/st-ember/chessbackend/api"
)

var ErrUnauthorized = errors.New("invalid username or token")

// checks token validity
func Authorization(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var token = r.Header.Get("Authorization")
		// var err error

		if token == "" {
			log.Error(ErrUnauthorized)
			api.RequestErrorHandler(w, ErrUnauthorized)
			return
		}
		// Todo: Implement jwt auth

		next.ServeHTTP(w, r)
	})
}
