package handlers

import (
	"net/http"
)


func Counter(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, World!"))
}