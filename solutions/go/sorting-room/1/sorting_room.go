package sorting

import (
	"fmt"
	"strconv"
)

// DescribeNumber should return a string describing the number.
func DescribeNumber(f float64) string {
	return fmt.Sprintf("This is the number %.1f", f)
}

type NumberBox interface {
	Number() int
}

// DescribeNumberBox should return a string describing the NumberBox.
func DescribeNumberBox(nb NumberBox) string {
	return fmt.Sprintf("This is a box containing the number %.1f", float64(nb.Number()))
}

type FancyNumber struct {
	n string
}

func (i FancyNumber) Value() string {
	return i.n
}

type FancyNumberBox interface {
	Value() string
}

// ExtractFancyNumber should return the integer value for a FancyNumber
// and 0 if any other FancyNumberBox is supplied.
func ExtractFancyNumber(fnb FancyNumberBox) int {
	f, ok := fnb.(FancyNumber)
	if ok {
		i, _ := strconv.Atoi(f.Value())
		return i
	}
	return 0
}

// DescribeFancyNumberBox should return a string describing the FancyNumberBox.
func DescribeFancyNumberBox(fnb FancyNumberBox) string {
	switch fnb.(type) {
	case FancyNumber:
		//i, _ := strconv.ParseFloat(f.n, 64)
		return fmt.Sprintf("This is a fancy box containing the number %.1f", float64(ExtractFancyNumber(fnb)))
	default:
		return "This is a fancy box containing the number 0.0"
	}
}

// DescribeAnything should return a string describing whatever it contains.
func DescribeAnything(i any) string {
	switch rel := i.(type) {
	case int:
		return DescribeNumber(float64(rel))
	case float64:
		return DescribeNumber(rel)
	case NumberBox:
		return DescribeNumberBox(rel)
	case FancyNumberBox:
		return DescribeFancyNumberBox(rel)
	default:
		return "Return to sender"
	}
}
