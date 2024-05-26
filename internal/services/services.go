package services

import (
	"encoding/gob"
	"log"
	"os"
	"time"
)


type CounterService struct {}

// Counter returns the number of requests made in the last 60 seconds.
func (c *CounterService) Counter(f string) (int, error){
	now := time.Now()
	cutoff := now.Add(-60 * time.Second)

	count := 1
	var requests []time.Time
	file, err := os.OpenFile(f, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		log.Println("Error opening file:", err)
		return 0, err
}
err = gob.NewDecoder(file).Decode(&requests)
if err != nil {
		requests = append(requests, now)
		err = gob.NewEncoder(file).Encode(requests)
		if err != nil {
			log.Println("Error encoding requests:", err)
			return 0, err
		}
		return 1, nil
}
	defer file.Close()
	for _, req := range requests {
		if req.After(cutoff) {
			count++
		}
	}
	// overwrite the file with the updated requests slice
	requests = append(requests, now)
	// fmt.Println(requests)

	// Truncate the file and move the file pointer to the beginning of the file
	err = file.Truncate(0)
	if err != nil {
			log.Println("Error truncating file:", err)
			return 0, err
	}
	_, err = file.Seek(0, 0)
	if err != nil {
			log.Println("Error seeking file:", err)
			return 0, err
	}

	err = gob.NewEncoder(file).Encode(requests)
	if err != nil {
			log.Println("Error encoding requests:", err)
			return 0, err
	}

	return count, nil
}


// SaveRequests saves the requests slice to a file.
// func (c *CounterService) SaveRequests(f string) error{
// 	file, err := os.OpenFile(f, os.O_WRONLY|os.O_CREATE, 0666)
// 	if err != nil {
// 		log.Println("Error opening file:", err)
// 		return err
// 	}
// 	err = gob.NewEncoder(file).Encode(requests)
// 	fmt.Println(requests)
// 	if err != nil {
// 		log.Println("Error encoding requests:", err)
// 		return err
// 	}
// 	return nil
// }

// LoadRequests loads the requests slice from a file.
// func (c *CounterService) LoadRequests(f string) ([]time.Time, error){
// 	file, err := os.Open(f)
// 	if err != nil {
// 		log.Println("Error opening file:", err)
// 		return nil, err
// 	}

// 	defer file.Close()

// 	err = gob.NewDecoder(file).Decode(&requests)
// 	if err != nil {
// 		log.Println("Error decoding requests:", err)
// 		return nil, err
// 	}
// 	return requests, nil
// }