package main

import (
	"net/http"
	"reflect"
	"testing"
)

func TestRoutes(t *testing.T) {
	mux := routes()

	if reflect.TypeOf(mux) != reflect.TypeOf(&http.ServeMux{}) {
		t.Errorf("Expected type to be *http.ServeMux, got %s", reflect.TypeOf(mux))
	}
}