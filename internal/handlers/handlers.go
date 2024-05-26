package handlers

import (
	"fmt"
	"net/http"
	"time"
)

var (
	requests []time.Time
)

func Counter(w http.ResponseWriter, r *http.Request) {
	now := time.Now()
	requests = append(requests, now)
	cutoff := now.Add(-60 * time.Second)

	count := 0
	for _, req := range requests {
		if req.After(cutoff) {
			count++
		}
	}
	fmt.Fprintf(w, "Number of requests in the last 60 seconds: %d\n", count)
}