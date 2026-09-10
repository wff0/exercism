package armstrongnumbers

import "math"

func IsNumber(n int) bool {
	sum := n
	var digits []int

	for n > 0 {
		digit := n % 10
		digits = append(digits, digit)
		n /= 10
	}

	for _, d := range digits {
		sum -= int(math.Pow(float64(d), float64(len(digits))))
	}
	return sum == 0
}
