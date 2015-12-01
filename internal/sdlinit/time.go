package sdlinit

import (
	"sync"
	"time"
)

var (
	initTimeMu sync.RWMutex // guards initTime

	// initTime is set to when SDL was initialized for syncing timestamps,
	// which are represented in milliseconds after SDL's initialization.
	initTime time.Time
)

func setInitTime(t time.Time) {
	initTimeMu.Lock()
	initTime = t
	initTimeMu.Unlock()
}

func getInitTime() time.Time {
	initTimeMu.RLock()
	t := initTime
	initTimeMu.RUnlock()
	return t
}

func init() {
	// If the user does not call one of the init functions, let's
	// at least give them a roughly sensible time.
	setInitTime(time.Now())
}

// TODO(eandre) fix wrapping of timestamp? (takes ~49 days)
func TimestampToTime(ts uint32) time.Time {
	initTime := getInitTime()

	// Timestamps are in milliseconds
	return initTime.Add(time.Duration(ts) * time.Millisecond)
}
