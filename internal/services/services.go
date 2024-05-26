package services

import "time"


type CounterService struct {}

var (
	requests []time.Time
)

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