// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package acronym should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package acronym

import (
	"regexp"
	"strings"
	"unicode"
)

// Abbreviate should have a comment documenting it.
func Abbreviate(s string) string {
	// Write some code here to pass the test suite.
	// Then remove all the stock comments.
	// They're here to help you get started but they only clutter a finished solution.
	// If you leave them in, reviewers may protest!
	reg := regexp.MustCompile(`[-\s_]+`)
	split := reg.Split(s, -1)
	sb := new(strings.Builder)
	for _, word := range split {
		if len(word) == 0 {
			continue
		}
		for _, r := range word {
			if !unicode.IsLetter(r) {
				continue
			}
			sb.WriteRune(unicode.ToUpper(r))
			break
		}
	}
	return sb.String()
}
