package handlers

import "github.com/takahiromitsui/go-server-counter/internal/config"


var app config.AppConfig

func setup(){
	app.FilePath = "requests.gob"
	repo := NewRepo(&app)
	NewHandlers(repo)
}