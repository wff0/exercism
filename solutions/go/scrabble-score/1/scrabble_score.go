package scrabblescore

import "strings"

func point(letter byte) int {
	switch letter {
	case 'a', 'e', 'i', 'o', 'u', 'l', 'n', 'r', 's', 't':
		return 1
	case 'd', 'g':
		return 2
	case 'b', 'c', 'm', 'p':
		return 3
	case 'f', 'h', 'v', 'w', 'y':
		return 4
	case 'k':
		return 5
	case 'j', 'x':
		return 8
	case 'q', 'z':
		return 10
	default:
		return 0
	}
}

func Score(word string) int {
	word = strings.ToLower(word)
	sum := 0
	for i := range word {
		sum += point(word[i])
	}
	return sum
}
