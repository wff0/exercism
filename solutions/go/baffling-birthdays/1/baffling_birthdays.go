package bafflingbirthdays

import (
	"math/rand/v2"
	"time"
)

func SharedBirthday(dates []time.Time) bool {
	set := make(map[string]struct{}, len(dates))
	for _, t := range dates {
		s := t.Format("01-02")
		if _, exist := set[s]; exist {
			return true
		}
		set[s] = struct{}{}
	}
	return false
}

func random() time.Time {
	var sec int64 = 0
	for range 10 {
		sec = sec*10 + rand.Int64N(10)
	}

	return time.Unix(sec, 0)
}

func IsLeapYear(year int) bool {
	return year%4 == 0 && year%100 != 0 || year%400 == 0
}

func RandomBirthdates(size int) []time.Time {
	res := make([]time.Time, 0, size)

	for range size {
		var t time.Time
		for t = random(); IsLeapYear(t.Year()); t = random() {

		}
		res = append(res, t)
	}
	return res
}

func EstimatedProbability(size int) float64 {
	if size == 1 {
		return 0
	}
	var res float64 = 1
	for i := 1; i < size; i++ {
		res *= float64(365-i) / 365
	}
	res = 1 - res
	return res * 100
}
