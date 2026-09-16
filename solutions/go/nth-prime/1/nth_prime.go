package nthprime

import "errors"

// Nth returns the nth prime number. An error must be returned if the nth prime number can't be calculated ('n' is equal or less than zero)
func Nth(n int) (int, error) {
	if n < 1 {
		return 0, errors.ErrUnsupported
	}
	var num = 1
	for count := 0; count < n; {
		num++
		if isPrime(num) {
			count++
		}
	}
	return num, nil
}

func isPrime(n int) bool {
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}
