package rotationalcipher

import (
	"strings"
	"unicode"
)

func RotationalCipher(plain string, shiftKey int) string {
	sb := new(strings.Builder)
	for _, r := range plain {
		var base rune
		if unicode.IsUpper(r) {
			base = 'A'
		} else if unicode.IsLower(r) {
			base = 'a'
		} else {
			sb.WriteRune(r)
			continue
		}
		i := int(r - base)
		i = (i + shiftKey) % 26
		sb.WriteRune(rune(int(base) + i))
	}
	return sb.String()
}
