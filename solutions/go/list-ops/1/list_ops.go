package listops

import "slices"

// IntList is an abstraction of a list of integers which we can define methods on
type IntList []int

func (s IntList) Foldl(fn func(int, int) int, initial int) int {
	for i := range s {
		initial = fn(initial, s[i])
	}
	return initial
}

func (s IntList) Foldr(fn func(int, int) int, initial int) int {
	for _, v := range slices.Backward(s) {
		initial = fn(v, initial)
	}
	return initial
}

func (s IntList) Filter(fn func(int) bool) IntList {
	result := make(IntList, 0, len(s))
	for _, v := range s {
		if fn(v) {
			result = append(result, v)
		}
	}
	return result
}

func (s IntList) Length() int {
	return len(s)
}

func (s IntList) Map(fn func(int) int) IntList {
	res := make(IntList, len(s))
	for i, v := range s {
		res[i] = fn(v)
	}
	return res
}

func (s IntList) Reverse() IntList {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}

func (s IntList) Append(lst IntList) IntList {
	return append(s, lst...)
}

func (s IntList) Concat(lists []IntList) IntList {
	l := len(s)
	for _, list := range lists {
		l += len(list)
	}
	res := make(IntList, 0, l)
	copy(res, s)
	for _, list := range lists {
		res = append(res, list...)
	}
	return res
}
