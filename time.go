package dataparse

import (
	"fmt"
	"time"
)

var ParseTimeFormats = []string{
	// most common formats
	time.RFC3339,
	time.RFC3339Nano,
	time.DateTime,
	time.DateOnly,
	time.TimeOnly,

	// likely more common than the remaining rfc formats
	time.ANSIC,
	time.UnixDate,

	// remaining rfc formats
	time.RFC822,
	time.RFC822Z,
	time.RFC850,
	time.RFC1123,
	time.RFC1123Z,

	// everything else
	time.Layout,
	time.RubyDate,
	time.Kitchen,
	time.Stamp,
	time.StampMilli,
	time.StampMicro,
	time.StampNano,
}

// ParseTime attempts to parse s as time utilizing all formats in
// ParseTimeFormats.
func ParseTime(s string) (time.Time, error) {
	for _, format := range ParseTimeFormats {
		t, err := time.Parse(format, s)
		if err == nil {
			return t, nil
		}
	}
	return time.Time{}, fmt.Errorf("dataparse: no format produced a valid time.Time")
}
