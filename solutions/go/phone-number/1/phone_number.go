package phonenumber

import (
	"errors"
	"fmt"
	"strings"
	"unicode"
)

func Number(phoneNumber string) (string, error) {
	phoneNumber = strings.NewReplacer(
		" ", "",
		".", "",
		"-", "",
		"+", "",
		"(", "",
		")", "",
	).Replace(phoneNumber)

	if len(phoneNumber) < 10 || len(phoneNumber) > 11 {
		return "", errors.ErrUnsupported
	}

	if len(phoneNumber) == 11 {
		if phoneNumber[0] != '1' {
			return "", errors.ErrUnsupported
		}
		phoneNumber = phoneNumber[1:]
	}

	if phoneNumber[0] < '2' || phoneNumber[3] < '2' {
		return "", errors.ErrUnsupported
	}

	for _, index := range []int{1, 2, 4, 5, 6, 7, 8, 9} {
		if !unicode.IsDigit(rune(phoneNumber[index])) {
			return "", errors.ErrUnsupported
		}
	}
	return phoneNumber, nil
}

func AreaCode(phoneNumber string) (string, error) {
	var err error
	phoneNumber, err = Number(phoneNumber)
	if err != nil {
		return "", err
	}
	return phoneNumber[:3], nil
}

func Format(phoneNumber string) (string, error) {
	var err error
	phoneNumber, err = Number(phoneNumber)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("(%s) %s-%s",
		phoneNumber[:3], phoneNumber[3:6], phoneNumber[6:]), nil
}
