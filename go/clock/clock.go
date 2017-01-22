// Package clock provides a clock which is independent from the date.
package clock

import "fmt"

const testVersion = 4
const minPerDay = 1440

// Clock represents a 24h clock
type Clock int

// New creates a new clock instance
func New(hour, minute int) Clock {
	return Clock(0).Add(hour*60 + minute)
}

func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c/60, c%60)
}

// Add moves the clock forward (backwards if negative) min minutes
func (c Clock) Add(min int) Clock {
	c = Clock((int(c) + min) % minPerDay)

	if c < 0 {
		c += minPerDay
	}
	return c
}
