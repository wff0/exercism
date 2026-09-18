package primefactors

func Factors(n int64) []int64 {
	res := make([]int64, 0)
	var prime int64 = 2
	for n > 1 {
		if n%prime == 0 {
			n /= prime
			res = append(res, prime)
		} else {
			for {
				prime++
				if isPrime(int(prime)) {
					break
				}
			}
		}
	}
	return res
}

func isPrime(n int) bool {
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}
