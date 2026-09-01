package affinecipher

import (
	"errors"
	"strings"
	"unicode"
)

func Encode(text string, a, b int) (string, error) {
	if gcd(a, 26) != 1 {
		return "", errors.ErrUnsupported
	}
	sb := &strings.Builder{}
	count := 0
	for _, r := range strings.ToLower(text) {
		if count != 0 && count%5 == 0 && unicode.IsLetter(r) {
			sb.WriteByte(' ')
		}
		if unicode.IsDigit(r) {
			sb.WriteRune(r)
			count++
		} else if unicode.IsLetter(r) {
			x := int(r - 'a')
			bt := byte(Ex(a, x, b, 26) + 'a')
			sb.WriteByte(bt)
			count++
		} else {
			continue
		}
	}
	return sb.String(), nil
}

func Decode(text string, a, b int) (string, error) {
	if gcd(a, 26) != 1 {
		return "", errors.ErrUnsupported
	}

	x := mmi(a, 26)
	sb := &strings.Builder{}
	for _, r := range strings.ToLower(text) {
		if unicode.IsDigit(r) {
			sb.WriteRune(r)
		} else if unicode.IsLetter(r) {
			y := int(r - 'a')
			bt := byte(Dy(x, y, b, 26) + 'a')
			sb.WriteByte(bt)
		} else {
			continue
		}
	}
	return sb.String(), nil
}

func Ex(a, i, b, m int) int {
	return (a*i + b) % m
}

func Dy(x, y, b, m int) int {
	return mod(x*(y-b), m)
}

func mod(a, m int) int {
	r := a % m
	if r < 0 {
		r += m
	}
	return r
}

func gcd(a, b int) int {
	if a < b {
		return gcd(b, a)
	}
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func mmi(a, m int) int {
	var i int
	for i = 1; (a*i)%m != 1; i++ {
	}
	return i
}
