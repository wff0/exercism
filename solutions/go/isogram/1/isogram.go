package isogram

import (
	"strings"
	"unicode"
)

func IsIsogram(word string) bool {
	countMap := make(map[rune]struct{}, len(word))
	for _, r := range strings.ToLower(word) {
		if !unicode.IsLetter(r) {
			continue
		}
		if _, ok := countMap[r]; ok {
			return false
		}
		countMap[r] = struct{}{}
	}
	return true
}
