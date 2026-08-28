package thefarm

import (
	"errors"
	"fmt"
)

// TODO: define the 'DivideFood' function
func DivideFood(cal FodderCalculator, cows int) (float64, error) {
	amount, err := cal.FodderAmount(cows)
	if err != nil {
		return 0, err
	}

	factor, err := cal.FatteningFactor()
	if err != nil {
		return 0, err
	}
	return amount * factor / float64(cows), nil
}

var errInvalidNumber = errors.New("invalid number of cows")

func ValidateInputAndDivideFood(cal FodderCalculator, cows int) (float64, error) {
	if cows <= 0 {
		return 0, errInvalidNumber
	}
	return DivideFood(cal, cows)
}

// TODO: define the 'ValidateNumberOfCows' function
type InvalidCowsError struct {
	Cows    int
	Message string
}

func (i *InvalidCowsError) Error() string {
	return fmt.Sprintf("%d cows are invalid: %s", i.Cows, i.Message)
}

func ValidateNumberOfCows(cows int) error {
	if cows < 0 {
		return &InvalidCowsError{
			Cows:    cows,
			Message: "there are no negative cows",
		}
	} else if cows == 0 {
		return &InvalidCowsError{
			Cows:    cows,
			Message: "no cows don't need food",
		}
	}
	return nil
}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
