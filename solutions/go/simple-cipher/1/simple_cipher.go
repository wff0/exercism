package cipher

import (
	"strings"
	"unicode"
)

// Define the shift and vigenere types here.
// Both types should satisfy the Cipher interface.

func NewCaesar() Cipher {
	return NewShift(3)
}

func NewShift(distance int) Cipher {
	if distance == 0 {
		return nil
	}
	if distance < -25 || distance > 25 {
		return nil
	}
	return shift{distance: distance}
}

type shift struct {
	distance int
}

func (c shift) Encode(input string) string {
	sb := new(strings.Builder)
	for _, r := range strings.ToLower(input) {
		if !unicode.IsLetter(r) {
			continue
		}
		base := int(r-'a') + c.distance
		if base >= 26 {
			base -= 26
		} else if base < 0 {
			base += 26
		}
		sb.WriteRune(rune(base + 'a'))
	}
	return sb.String()
}

func (c shift) Decode(input string) string {
	sb := new(strings.Builder)
	for _, r := range input {
		if !unicode.IsLetter(r) {
			continue
		}
		base := int(r-'a') - c.distance
		if base >= 26 {
			base -= 26
		} else if base < 0 {
			base += 26
		}
		sb.WriteRune(rune(base + 'a'))
	}
	return sb.String()
}

func NewVigenere(key string) Cipher {
	if len(key) == 0 {
		return nil
	}
	flag := 0
	for _, r := range key {
		if !unicode.IsLetter(r) {
			return nil
		}
		if unicode.ToLower(r) != r {
			return nil
		}

		flag += int(r - 'a')
	}
	if flag == 0 {
		return nil
	}
	return vigenere{key: key}
}

type vigenere struct {
	key string
}

func (v vigenere) Encode(input string) string {
	index := 0
	sb := new(strings.Builder)
	for _, r := range strings.ToLower(input) {
		if !unicode.IsLetter(r) {
			continue
		}
		index %= len(v.key)
		base := int(r-'a') + int(v.key[index]-'a')
		if base >= 26 {
			base -= 26
		} else if base < 0 {
			base += 26
		}
		sb.WriteRune(rune(base + 'a'))
		index++
	}
	return sb.String()
}

func (v vigenere) Decode(input string) string {
	index := 0
	sb := new(strings.Builder)
	for _, r := range input {
		if !unicode.IsLetter(r) {
			continue
		}
		index %= len(v.key)
		base := int(r-'a') - int(v.key[index]-'a')
		if base >= 26 {
			base -= 26
		} else if base < 0 {
			base += 26
		}
		sb.WriteRune(rune(base + 'a'))
		index++
	}
	return sb.String()
}
