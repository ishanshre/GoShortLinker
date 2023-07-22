package main

import (
	"log"
	"net/http"
	"os"

	"github.com/ishanshre/GoShortLinker/internals/database"
	"github.com/ishanshre/GoShortLinker/internals/handlers"
	"github.com/ishanshre/GoShortLinker/internals/routers"
	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatalf("error in loading environment files: %v", err)
	}
}

func main() {
	db, err := database.NewConnection(os.Getenv("dsn"))
	if err != nil {
		log.Fatalf("error in connecting to database: %v", err)
	}
	log.Println("Connected to Database")
	h := handlers.NewHandler(db)
	r := routers.Router(h)
	srv := http.Server{
		Addr:    ":8000",
		Handler: r,
	}
	log.Println("Starting server at :8000")
	if err := srv.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
