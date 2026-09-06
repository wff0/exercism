package reversestring

import (
	"slices"
	"strings"
)

func Reverse(input string) string {
	sb := new(strings.Builder)
	rns := []rune(input)
	for _, rn := range slices.Backward(rns) {
		sb.WriteRune(rn)
	}
	return sb.String()
}
