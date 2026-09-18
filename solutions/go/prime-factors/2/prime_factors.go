package primefactors

func Factors(n int64) []int64 {
	res := make([]int64, 0)
	var prime int64 = 2
	for n > 1 {
		for n%prime == 0 {
			n /= prime
			res = append(res, prime)
		}
		prime++
	}
	return res
}
