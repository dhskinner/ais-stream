package atlas

import (
	"time"
)

// Calculate the time to go to the next event, based on a given interval, and offset. For example:
//  1. to generate an event once per minute, 30 seconds after the top of each minute:
//     timeToGo( interval = 60 seconds, offset = 30 seconds)
//  2. to generate an event once per day, at 12:15pm AEST:
//     timeToGo( interval = 24 hours, offset = 2 hours 15 minutes)
func timeToGo(interval time.Duration, offset time.Duration) time.Duration {
	timePrev := time.Unix(0, int64((time.Now().UnixNano()-int64(offset))/int64(interval))*int64(interval)+int64(offset))
	timeNext := timePrev.Add(interval)
	return time.Until(timeNext)
}
