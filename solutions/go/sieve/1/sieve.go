package sieve

func Sieve(limit int) []int {
	nums := make([]int, 0, limit-1)
	mark := make(map[int]bool, limit-1)
	for i := 2; i <= limit; i++ {
		nums = append(nums, i)
		mark[i] = false
	}

	res := make([]int, 0)
	for _, n := range nums {
		if mark[n] {
			continue
		}
		if isPrime(n) {
			res = append(res, n)
		}
		for i := 2; i*n <= limit; i++ {
			if isPrime(i * n) {
				mark[i*n] = true
			}
		}
	}
	return res
}

func isPrime(n int) bool {
	for i := 2; i*i < n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}
