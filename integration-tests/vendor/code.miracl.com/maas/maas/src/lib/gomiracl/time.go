package gomiracl

import "time"

const secondsInADay = 24 * 60 * 60

// EpochDate returns epoch time in days.
func EpochDate() int {
	return int(time.Now().Unix() / secondsInADay)
}
