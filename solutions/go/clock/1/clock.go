package clock

import "fmt"

// Define the Clock type here.
type Clock struct {
	m int
}

func New(h, m int) Clock {
	m += h * 60
	for m >= 24*60 {
		m -= 24 * 60
	}
	for m < 0 {
		m += 24 * 60
	}
	return Clock{
		m: m,
	}
}

func (c Clock) Add(m int) Clock {
	return New(0, c.m+m)
}

func (c Clock) Subtract(m int) Clock {
	return New(0, c.m-m)
}

func (c Clock) String() string {
	for c.m < 0 {
		c.m += 24 * 60
	}
	h := c.m / 60 % 24
	m := c.m % 60

	return fmt.Sprintf("%02d:%02d", h, m)
}
