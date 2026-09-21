package squareroot

func SquareRoot(number int) (int, error) {
	if number == 1 {
		return 1, nil
	}
	left, right := 1, number
	for left < right {
		mid := left + (right-left)>>1
		if mid*mid < number {
			left = mid + 1
		} else if mid*mid > number {
			right = mid - 1
		} else {
			return mid, nil
		}
	}
	return left, nil
}
