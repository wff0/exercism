package wordcount

import (
	"regexp"
	"strings"
)

type Frequency map[string]int

func WordCount(phrase string) Frequency {
	phrase = strings.ToLower(phrase)
	reg := regexp.MustCompile(`\w+('(\w)+)?`)
	words := reg.FindAllString(phrase, -1)

	freq := make(Frequency)
	for _, word := range words {
		freq[word]++
	}
	return freq
}
