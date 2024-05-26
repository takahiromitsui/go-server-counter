package main

import (
	"fmt"
	"net/http"
)

const port = ":8080"

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc(("GET /counter"), func (w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, World!"))
	})
	fmt.Println("Server is running on port", port)
	http.ListenAndServe(port, mux)
}