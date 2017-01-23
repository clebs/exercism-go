// Package gigasecond provides logic to manipulate gigasecond time units and make them compatible with time.Time.
package gigasecond

import "time"

const testVersion = 4

const gigaSecond = time.Second * 1e9

// API function.  It uses a type from the Go standard library.
func AddGigasecond(t time.Time) time.Time {
	return t.Add(gigaSecond)
}
