package sumofmultiples

func SumMultiples(limit int, divisors ...int) int {
	if len(divisors) == 0 {
		return 0
	}
	set := map[int]struct{}{}
	for _, v := range divisors {
		if v == 0 {
			continue
		}
		count := 1
		for v*count < limit {
			set[v*count] = struct{}{}
			count++
		}
	}

	sum := 0
	for k := range set {
		sum += k
	}
	return sum
}
