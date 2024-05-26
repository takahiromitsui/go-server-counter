package services

import (
    "os"
    "testing"
    "time"
)

func TestLoadRequests(t *testing.T) {
    // Create a temporary file for testing
    file, err := os.CreateTemp("", "test")
    if err != nil {
        t.Fatal(err)
    }
    defer os.Remove(file.Name())

    cs := &CounterService{}
    requests, err := cs.loadRequests(file.Name())
    if err != nil {
        t.Fatal(err)
    }
    if len(requests) != 0 {
        t.Errorf("Expected requests to be empty, got %v", requests)
    }
}

func TestSaveRequests(t *testing.T) {
    // Create a temporary file for testing
    file, err := os.CreateTemp("", "test")
    if err != nil {
        t.Fatal(err)
    }
    defer os.Remove(file.Name())

    cs := &CounterService{}

    // Test that saveRequests saves the requests to the file
    err = cs.saveRequests(file.Name(), []time.Time{time.Now()})
    if err != nil {
        t.Fatal(err)
    }

    // Load the requests from the file and check that they match the saved requests
    requests, err := cs.loadRequests(file.Name())
    if err != nil {
        t.Fatal(err)
    }
    if len(requests) != 1 {
        t.Errorf("Expected one request, got %v", requests)
    }
}

func TestCountRequests(t *testing.T) {
    cs := &CounterService{}

    // Test that countRequests returns the correct count
		// The first argument is the list of requests, the second argument is the cutoff time
    count := cs.countRequests([]time.Time{time.Now(), time.Now().Add(-600 * time.Second)}, time.Now().Add(-60 * time.Second))
    if count != 1 {
        t.Errorf("Expected count to be 1, got %d", count)
    }
		// Test that countRequests returns the correct count
		count = cs.countRequests([]time.Time{time.Now(), time.Now().Add(-50 * time.Second)}, time.Now().Add(-60 * time.Second))
		if count != 2 {
				t.Errorf("Expected count to be 2, got %d", count)
		}
}