package differenceofsquares

func SquareOfSum(n int) int {
	sum := (1 + n) * n / 2
	//for i := 1; i <= n; i++ {
	//	sum += i
	//}
	return sum * sum
}

func SumOfSquares(n int) int {
	sum := n * (n + 1) * (2*n + 1) / 6
	//for i := 1; i <= n; i++ {
	//	sum += i * i
	//}
	return sum
}

func Difference(n int) int {
	diff := SumOfSquares(n) - SquareOfSum(n)
	if diff < 0 {
		return -diff
	}
	return diff
}
