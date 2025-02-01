package auth

import (
	"encoding/json"
	"net/http"

	log "github.com/sirupsen/logrus"

	"github.com/st-ember/chessbackend/api"
	"github.com/st-ember/chessbackend/internal/db"
	"github.com/st-ember/chessbackend/internal/tools"
)

func Login(w http.ResponseWriter, r *http.Request) {
	var params LoginParams
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&params)
	if err != nil {
		log.Error("cannot decode request body", err)
		api.InternalErrorHandler(w)
		return
	}

	username := params.Username
	password := params.Password

	dbPwd, err := db.RetrievePassword(username)
	if err != nil {
		log.Error("cannot retrieve password", err)
		api.InternalErrorHandler(w)
		return
	}

	encPwd, err := tools.EncryptAESG(password)
	if encPwd != dbPwd {
		log.Error("invalid password")
		api.InternalErrorHandler(w)
		return
	}

	accessToken, refreshToken, refreshClaims, err := tools.GenerateTokens(username)
	if err != nil {
		log.Error("cannot generate tokens", err)
		api.InternalErrorHandler(w)
		return
	}

	response := api.BaseResponse{
		Code:    http.StatusOK,
		Success: true,
		Message: "Login Successful",
		Data: map[string]interface{}{
			"username": username,
			"token":    accessToken,
		},
	}

	cookie := &http.Cookie{
		Name:     "refresh_token",
		Value:    refreshToken,
		Expires:  refreshClaims.ExpiresAt.Time,
		HttpOnly: true,
		Secure:   false,
		SameSite: http.SameSiteLaxMode,
	}

	http.SetCookie(w, cookie)

	w.Header().Set("Content-Type", "application/json")

	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		log.Error("cannot encode response", err)
		api.InternalErrorHandler(w)
		return
	}
}
