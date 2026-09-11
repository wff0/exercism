// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package bob should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package bob

import (
	"strings"
)

import "unicode"

// Hey should have a comment documenting it.
func Hey(remark string) string {
	// Write some code here to pass the test suite.
	// Then remove all the stock comments.
	// They're here to help you get started but they only clutter a finished solution.
	// If you leave them in, reviewers may protest!
	remark = strings.TrimSpace(remark)
	if IsSilent(remark) {
		return "Fine. Be that way!"
	}
	isQuestion := IsQuestion(remark)
	isYell := IsYell(remark)
	if isQuestion && isYell {
		return "Calm down, I know what I'm doing!"
	}
	if isQuestion {
		return "Sure."
	}
	if isYell {
		return "Whoa, chill out!"
	}
	return "Whatever."
}

func IsQuestion(remark string) bool {
	return remark[len(remark)-1] == '?'
}

func IsYell(remark string) bool {
	rs := []rune(remark)
	hasLetter := false
	for _, r := range rs {
		if unicode.IsLetter(r) {
			hasLetter = true
		}
		if unicode.IsLower(r) {
			return false
		}
	}
	return hasLetter
}

func IsSilent(remark string) bool {
	return len(remark) == 0
}
