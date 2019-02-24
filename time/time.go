/*
Package T provides the protosocial time.

This root package defines a universal time;
for more pragmatic, planetary-specific times, see the subpackages.
*/
package T // import "gitlab.com/hall/measure/time"

import (
	"strconv"
	_time "time"
)

// Conv converts a time to prototime.
func Conv(t _time.Time) string {
	t = t.UTC()
	// Shift year 0 to the beginning of the Holocene period.
	year := strconv.FormatInt(int64(t.Year()+11700), 12)
	// The twelve months are numbered from 0 to b.
	month := strconv.FormatInt(int64(t.Month())-1, 12)
	// The first day of a month is 0.
	day := strconv.FormatInt(int64(t.Day())-1, 12)
	// Twelve hours in a day (0 to b).
	hour := strconv.FormatInt(int64((t.Hour())/2), 12)
	// Twelve minutes in an hour.
	minute := strconv.FormatInt(int64((t.Minute())/5), 12)
	// Twelve seconds in a minute.
	second := strconv.FormatInt(int64((t.Second())/5), 12)
	//time := strconv.FormatInt(int64(hour+minute), 12)

	return year + "." + month + "|" + day + "." + hour + minute + second
}

// Now returns the current Time.
func Now() string {
	return Conv(_time.Now().UTC())
}
