package romannumerals

import (
	"errors"
	"strings"
)

type valueSymbol struct {
	symbol string
	value  int
}

var valueSymbols = []valueSymbol{
	{
		symbol: "M",
		value:  1000,
	},
	{
		symbol: "CM",
		value:  900,
	},
	{
		symbol: "D",
		value:  500,
	},
	{
		symbol: "CD",
		value:  400,
	},
	{
		symbol: "C",
		value:  100,
	},
	{
		symbol: "XC",
		value:  90,
	},
	{
		symbol: "L",
		value:  50,
	},
	{
		symbol: "XL",
		value:  40,
	},
	{
		symbol: "X",
		value:  10,
	},
	{
		symbol: "IX",
		value:  9,
	},
	{
		symbol: "V",
		value:  5,
	},
	{
		symbol: "IV",
		value:  4,
	},
	{
		symbol: "I",
		value:  1,
	},
}

func ToRomanNumeral(input int) (string, error) {
	if input < 1 || input > 3999 {
		return "", errors.ErrUnsupported
	}
	sb := new(strings.Builder)
	for _, v := range valueSymbols {
		for input >= v.value {
			sb.WriteString(v.symbol)
			input -= v.value
		}
		if input == 0 {
			break
		}
	}
	return sb.String(), nil
}
