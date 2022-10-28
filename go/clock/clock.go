package clock

import "fmt"

// Define the Clock type here.
type Clock struct {
	minute, hour int
}

func New(h, m int) Clock {
	minutes := m % 60
	hours := (h + (m / 60)) % 24
	if minutes < 0 {
		hours--
		minutes += 60
	}
	if hours < 0 {
		hours += 24
	}
	return Clock{minutes, hours}
}

func (c Clock) Add(m int) Clock {
	c.minute += m
	return New(c.hour, c.minute)
}

func (c Clock) Subtract(m int) Clock {
	c.minute -= m
	return New(c.hour, c.minute)
}

func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.hour, c.minute)
}
