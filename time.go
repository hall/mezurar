/*
Package T provides the protosocial time.

The two main units of measurement are the day and the year.
This is a consequence of their convenience;
both are periods with useful frequencies.

  1 year = 1 revolution of the Earth around the Sun
  1 day  = 1 the rotation of the Earth on its axis

These definitions correspond to their contemporary usage (note that they are closely tied to the planet Earth).
This package departs convention by removing the arbitrary complexity of the remaining units (minutes, seconds, . . . );

Unfortunately, the frequencies are irreconcilable;
that is, we cannot evenly divide a year into whole days.
This is the source of the greatest complexity in this system of measurement.
*/
package T // import "gitlab.com/protosocial/time"

import (
	"strconv"
	_time "time"
)

// Time is the standard protosocial time.
type Time struct {
	year string
	day  string
}

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
