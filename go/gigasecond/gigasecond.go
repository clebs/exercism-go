// Package gigasecond provides logic to manipulate gigasecond time units and make them compatible with time.Time.
package gigasecond

import "time"

const testVersion = 4

const gigaSecond = time.Second * 1e9

// AddGigasecond adds a gigasecond to t and returns a copy of it.
func AddGigasecond(t time.Time) time.Time {
	return t.Add(gigaSecond)
}
