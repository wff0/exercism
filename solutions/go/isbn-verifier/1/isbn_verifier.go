package isbnverifier

import (
	"strings"
	"unicode"
)

func IsValidISBN(isbn string) bool {
	sum := 0
	count := 10
	isbn = strings.ReplaceAll(isbn, "-", "")
	if len(isbn) != 10 {
		return false
	}
	for i, r := range isbn {
		if r == 'X' {
			if i != len(isbn)-1 {
				return false
			}
			sum += 10 * count
			count--
		} else if unicode.IsDigit(r) {
			sum += int(r-'0') * count
			count--
		} else {
			return false
		}
	}
	return sum%11 == 0
}
