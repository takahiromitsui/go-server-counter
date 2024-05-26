package services

import (
	"encoding/gob"
	"io"
	"log"
	"os"
	"time"
)


type CounterService struct {}

// Counter returns the number of requests made in the last 60 seconds.
func (c *CounterService) Counter(f string) (int, error){
	now := time.Now()
	cutoff := now.Add(-60 * time.Second)

	requests, err := c.loadRequests(f)
	if err != nil {
			return 0, err
	}

	count := c.countRequests(requests, cutoff) // before appending the current request

	requests = append(requests, now)
	err = c.saveRequests(f, requests)
	if err != nil {
			return 0, err
	}

	return count+1, nil
}


// loadRequests loads the requests from the file.
func (c *CounterService) loadRequests(f string) ([]time.Time, error) {
	var requests []time.Time
	file, err := os.OpenFile(f, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
			log.Println("Error opening file:", err)
			return nil, err
	}
	defer file.Close()

	err = gob.NewDecoder(file).Decode(&requests)
	if err != nil && err != io.EOF {
			log.Println("Error decoding requests:", err)
			return nil, err
	}

	return requests, nil
}

// saveRequests saves the requests to the file.
func (c *CounterService) saveRequests(f string, requests []time.Time) error {
    file, err := os.OpenFile(f, os.O_RDWR|os.O_CREATE, 0666)
    if err != nil {
        log.Println("Error opening file:", err)
        return err
    }
    defer file.Close()

    err = file.Truncate(0)
    if err != nil {
        log.Println("Error truncating file:", err)
        return err
    }
    _, err = file.Seek(0, 0)
    if err != nil {
        log.Println("Error seeking file:", err)
        return err
    }

    err = gob.NewEncoder(file).Encode(requests)
    if err != nil {
        log.Println("Error encoding requests:", err)
        return err
    }

    return nil
}

// countRequests counts the number of requests made after the cutoff time.
func (c *CounterService) countRequests(requests []time.Time, cutoff time.Time) int {
    count := 0
    for _, req := range requests {
        if req.After(cutoff) {
            count++
        }
    }
    return count
}