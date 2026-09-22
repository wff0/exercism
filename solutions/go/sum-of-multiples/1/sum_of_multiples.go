package sumofmultiples

func SumMultiples(limit int, divisors ...int) int {
	if len(divisors) == 0 {
		return 0
	}
	index := 0
	for _, d := range divisors {
		if d != 0 {
			divisors[index] = d
			index++
		}
	}
	divisors = divisors[:index]
	if len(divisors) == 0 {
		return 0
	}
	set := map[int]struct{}{}
	for i := 1; ; i++ {
		hasValue := false
		for _, n := range divisors {
			if n*i < limit {
				set[n*i] = struct{}{}
				hasValue = true
			}
		}
		if !hasValue {
			break
		}
	}

	sum := 0
	for k := range set {
		sum += k
	}
	return sum
}
