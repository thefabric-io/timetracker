package timetracker

import (
	"log"
	"time"
)

func Track(start time.Time) time.Duration {
	return time.Since(start)
}

func Log(start time.Time, label string) {
	log.Printf("%s took %s", label, Track(start))
}
