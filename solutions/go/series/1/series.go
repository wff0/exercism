package series

func All(n int, s string) []string {
	if len(s) < n || n <= 0 {
		return nil
	}

	res := make([]string, 0)
	left, right := 0, 0
	for right < len(s) {
		if right-left+1 < n {
			right++
		} else {
			res = append(res, s[left:right+1])
			left++
			right++
		}
	}
	return res
}

func UnsafeFirst(n int, s string) string {
	if len(s) < n || n <= 0 {
		return ""
	}
	return s[0:n]
}

func First(n int, s string) (string, bool) {
	if len(s) < n || n <= 0 {
		return "", false
	}
	return s[0:n], true
}
