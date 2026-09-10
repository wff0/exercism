package anagram

import (
	"strings"
)

func Detect(subject string, candidates []string) []string {
	res := make([]string, 0, len(candidates))
	sMap := map[rune]int{}
	for _, r := range strings.ToLower(subject) {
		sMap[r]++
	}

	for _, candidate := range candidates {
		if strings.EqualFold(subject, candidate) || len(subject) != len(candidate) {
			continue
		}
		tMap := map[rune]int{}
		hasDiff := false
		for _, r := range strings.ToLower(candidate) {
			if _, ok := sMap[r]; !ok {
				hasDiff = true
				continue
			}
			tMap[r]++
		}
		if hasDiff {
			continue
		}

		count := 0
		for r, v := range tMap {
			if sMap[r] != v {
				continue
			}
			count++
		}
		if count == len(sMap) {
			res = append(res, strings.Clone(candidate))
		}
	}
	return res
}
