package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/takahiromitsui/go-server-counter/internal/config"
	"github.com/takahiromitsui/go-server-counter/internal/services"
)

type Repository struct {
	App *config.AppConfig
}

var Repo *Repository

// NewRepo creates a new Repository.
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{App: a}
}

// NewHandlers sets the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

type CounterResponse struct {
	Count int `json:"count"`
}

func (m *Repository)Counter(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(map[string]string{"error": "method not allowed"})
	}
	counterService := &services.CounterService{}
	count, err := counterService.Counter(m.App.FilePath)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "internal server error"})
		return
	}
	w.Header().Set("Content-Type", "application/json")
	resp := CounterResponse{Count: count}
	out, err := json.Marshal(resp)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "internal server error"})
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(out)
}