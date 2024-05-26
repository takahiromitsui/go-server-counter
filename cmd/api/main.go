package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/takahiromitsui/go-server-counter/internal/config"
	"github.com/takahiromitsui/go-server-counter/internal/handlers"
)

const port = ":8080"
var app config.AppConfig

func main() {
	app.FilePath = "requests.gob"
	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)
	file , err := os.Open(app.FilePath)
	if err != nil {
		log.Println("File not found:", err)
		log.Println("Creating new file...")
		file, err = os.Create(app.FilePath)
		if err != nil {
			log.Println("Error creating file:", err)
		}
	}
	defer file.Close()
	fmt.Println("Server is running on port", port)
	srv := &http.Server{
		Addr: port,
		Handler: routes(),
	}
	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}