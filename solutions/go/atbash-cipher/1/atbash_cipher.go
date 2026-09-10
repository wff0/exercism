package atbashcipher

import (
	"strings"
	"unicode"
)

func Atbash(s string) string {
	sb := new(strings.Builder)
	count := 0
	s = strings.ToLower(s)
	for _, r := range s {
		if unicode.IsDigit(r) {
			sb.WriteRune(r)
			count++
		} else if unicode.IsLetter(r) {
			if count != 0 && count%5 == 0 {
				sb.WriteByte(' ')
			}
			sb.WriteByte('z' - byte(r) + 'a')
			count++
		} else {
			continue
		}
	}
	return sb.String()
}
