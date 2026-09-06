package pangram

import "strings"

func IsPangram(input string) bool {
	input = strings.ToLower(input)
	count := [26]bool{}
	for i := range input {
		index := input[i] - 'a'
		if !(index >= 0 && index < 26) {
			continue
		}
		count[input[i]-'a'] = true
	}
	for _, b := range count {
		if !b {
			return false
		}
	}
	return true
}
