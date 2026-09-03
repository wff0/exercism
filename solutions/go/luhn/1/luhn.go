package luhn

import (
	"strings"
	"unicode"
)

func Valid(id string) bool {
	id = strings.ReplaceAll(id, " ", "")
	if len(id) <= 1 {
		return false
	}
	sum := 0
	count := 0

	for i := len(id) - 1; i >= 0; i-- {
		if unicode.IsDigit(rune(id[i])) {
			count++
			num := int(id[i] - '0')
			if count%2 == 0 {
				num *= 2
				if num > 9 {
					num -= 9
				}
			}
			sum += num
		} else {
			return false
		}
	}
	return sum%10 == 0
}
