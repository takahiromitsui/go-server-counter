package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCounter(t *testing.T) {
	setup()
	// Invalid method
	req, err := http.NewRequest("POST", "/counter", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(Repo.Counter)
	handler.ServeHTTP(rr, req)
	if rr.Code != http.StatusMethodNotAllowed {
		t.Errorf("Expected status 405, got %d", http.StatusMethodNotAllowed)
	}
	// Valid method
	req, err = http.NewRequest("GET", "/counter", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr = httptest.NewRecorder()
	handler.ServeHTTP(rr, req)
	if rr.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %d", rr.Code)
	}
	if rr.Body.String() == `{"count":0}` {
    t.Errorf("Expected count to be at least 1, got %s", rr.Body.String())
	}
}