package auth

import (
	"encoding/json"
	"net/http"

	log "github.com/sirupsen/logrus"

	"github.com/st-ember/chessbackend/api"
	"github.com/st-ember/chessbackend/internal/db"
	"github.com/st-ember/chessbackend/internal/tools"
)

func Signup(w http.ResponseWriter, r *http.Request) {
	var params SignupParams
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode((&params))
	if err != nil {
		log.Error("cannot decode request body", err)
		api.InternalErrorHandler(w)
		return
	}

	err = db.CreateUser(params.Username, params.Password, params.InitElo)
	if err != nil {
		log.Error("error creating user", err)
		api.InternalErrorHandler(w)
		return
	}

}
