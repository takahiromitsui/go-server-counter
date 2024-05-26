package services

import (
	"encoding/gob"
	"log"
	"os"
	"time"
)


type CounterService struct {}

var (
	requests []time.Time
)

// Counter returns the number of requests made in the last 60 seconds.
func (c *CounterService) Counter() int{
	now := time.Now()
	requests = append(requests, now)
	cutoff := now.Add(-60 * time.Second)

	count := 0
	for _, req := range requests {
		if req.After(cutoff) {
			count++
		}
	}
	return count
}

// SaveRequests saves the requests slice to a file.
func (c *CounterService) SaveRequests(f string) error{
	file, err := os.OpenFile(f, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		log.Println("Error opening file:", err)
		return err
	}
	err = gob.NewEncoder(file).Encode(requests)
	if err != nil {
		log.Println("Error encoding requests:", err)
		return err
	}
	return nil
}