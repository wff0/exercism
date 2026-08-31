package raindrops

import (
	"strconv"
	"strings"
)

func Convert(number int) string {
	res := &strings.Builder{}
	if number%3 == 0 {
		res.WriteString("Pling")
	}
	if number%5 == 0 {
		res.WriteString("Plang")
	}
	if number%7 == 0 {
		res.WriteString("Plong")
	}
	if res.Len() == 0 {
		return strconv.Itoa(number)
	}
	return res.String()
}
