package lineup

import "fmt"

const tpl = "%s, you are the %d%s customer we serve today. Thank you!"

func Format(name string, number int) string {
	last := number % 10
	last2 := number % 100

	var suffix string
	if last == 1 && last2 != 11 {
		suffix = "st"
	} else if last == 2 && last2 != 12 {
		suffix = "nd"
	} else if last == 3 && last2 != 13 {
		suffix = "rd"
	} else {
		suffix = "th"
	}
	return fmt.Sprintf(tpl, name, number, suffix)
}
