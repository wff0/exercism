package microblog

import (
	"strings"
	"unicode/utf8"
)

func Truncate(phrase string) string {
	n := utf8.RuneCountInString(phrase)
	if n <= 5 {
		return phrase
	}
	sb := new(strings.Builder)
	count := 0
	for count < 5 {
		rs, size := utf8.DecodeRuneInString(phrase)
		sb.WriteRune(rs)
		phrase = phrase[size:]
		count++
	}
	return sb.String()
}
