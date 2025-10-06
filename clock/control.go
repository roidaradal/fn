package clock

import "time"

// Run infinite loop; used at the end of main with goroutines
func RunForever() {
	for {
		select {}
	}
}

// Pauses the current goroutine for given duration,
// Takes into account starting time and adjusts the pause duration
func Sleep(pause time.Duration, start time.Time) {
	sleep := pause - TimeNow().Sub(start)
	time.Sleep(sleep)
}
