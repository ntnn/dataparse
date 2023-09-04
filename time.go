package dataparse

import (
	"fmt"
	"time"
)

// ParseTimeFormats are the various formats ParseTime and its consumers
// utilize to attempt to parse timestamps.
var ParseTimeFormats = []string{
	// Most common formats
	time.RFC3339,
	time.RFC3339Nano,
	time.DateTime,
	time.DateOnly,
	time.TimeOnly,

	// Likely more common than the remaining rfc formats
	time.ANSIC,
	time.UnixDate,

	// Remaining rfc formats
	time.RFC822,
	time.RFC822Z,
	time.RFC850,
	time.RFC1123,
	time.RFC1123Z,

	// Everything else from stdlib
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
