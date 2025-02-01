package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
	"github.com/st-ember/chessbackend/internal/db"
	"github.com/st-ember/chessbackend/internal/handlers"
)

func main() {
	log.SetReportCaller(true)
	var r *chi.Mux = chi.NewRouter()
	handlers.Handler(r)

	fmt.Println("Starting server")

	err := http.ListenAndServe("localhost:8000", r)
	if err != nil {
		log.Error(err)
	}

	fmt.Println("Connecting to Postgres")

	if err = godotenv.Load(); err != nil {
		log.Error("No .env file found")
	}

	db.InitDB(os.Getenv("CONNSTR"))
	defer db.CloseDB()
}
