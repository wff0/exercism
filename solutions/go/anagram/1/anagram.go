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
		if strings.EqualFold(subject, candidate) {
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

func isAnagram(subject [26]int, candidate string) bool {
	//if strings.EqualFold(subject, candidate) {
	//	return false
	//}
	//sArr := [26]int{}
	//for i := range subject {
	//	sArr[subject[i] - 'a']++
	//}
	tArr := [26]int{}
	for i := range candidate {
		tArr[candidate[i]-'a']++
	}
	return subject == tArr
}
