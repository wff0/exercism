package largestseriesproduct

import (
	"errors"
	"unicode"
)

func LargestSeriesProduct(digits string, span int) (int64, error) {
	if span < 0 || len(digits) < span {
		return 0, errors.ErrUnsupported
	}
	nums := make([]int64, 0, span)
	var res int64 = 0
	for i := 0; i < len(digits); i++ {
		if !unicode.IsDigit(rune(digits[i])) {
			return 0, errors.ErrUnsupported
		}
		if i < span {
			nums = append(nums, int64(digits[i]-'0'))
		} else {
			nums = nums[1:]
			nums = append(nums, int64(digits[i]-'0'))
		}
		if len(nums) == span {
			res = max(res, mul(nums))
		}
	}
	return res, nil
}

func mul(nums []int64) int64 {
	var res int64 = 1
	for _, n := range nums {
		res *= n
	}
	return res
}
