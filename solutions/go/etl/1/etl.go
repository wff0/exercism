package etl

import (
	"strings"
)

func Transform(in map[int][]string) map[string]int {
	res := map[string]int{}
	for val, ss := range in {
		for _, s := range ss {
			res[strings.ToLower(s)] = val
		}
	}

	return res
}
