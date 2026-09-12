package clock

import "fmt"

const (
	dayMinutes = 1440
	//hourMinutes = 60
)

// Define the Clock type here.
type Clock struct {
	m int
}

func New(h, m int) Clock {
	m = (h*60 + m) % dayMinutes
	if m < 0 {
		m += dayMinutes
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
	h := c.m / 60 % 24
	m := c.m % 60

	return fmt.Sprintf("%02d:%02d", h, m)
}
