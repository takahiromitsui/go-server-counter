package main

import (
	"fmt"
	"log"
	"net/http"
)

const port = ":8080"

func main() {
	fmt.Println("Server is running on port", port)
	srv := &http.Server{
		Addr: port,
		Handler: routes(),
	}
	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}