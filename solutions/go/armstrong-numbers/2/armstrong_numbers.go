package armstrongnumbers

func IsNumber(n int) bool {
	sum := n
	var digits []int

	for n > 0 {
		digit := n % 10
		digits = append(digits, digit)
		n /= 10
	}

	for _, d := range digits {
		//sum -= int(math.Pow(float64(d), float64(len(digits))))
		sum -= quickPow(d, len(digits))
	}
	return sum == 0
}

func quickPow(a, b int) int {
	res := 1
	for b > 0 {
		if b&1 == 1 {
			res *= a
		}
		a *= a
		b >>= 1
	}
	return res
}
