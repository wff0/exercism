package parsinglogfiles

import (
	"fmt"
	"regexp"
)

func IsValidLine(text string) bool {
	reg := regexp.MustCompile(`^\[(TRC|DBG|INF|WRN|ERR|FTL)\]`)

	return reg.MatchString(text)
}

func SplitLogLine(text string) []string {
	reg := regexp.MustCompile(`<[~*=-]*>`)
	return reg.Split(text, -1)
}

func CountQuotedPasswords(lines []string) int {
	count := 0
	reg := regexp.MustCompile(`".*(?i)password.*"`)
	for _, line := range lines {
		if reg.MatchString(line) {
			count++
		}
	}
	return count
}

func RemoveEndOfLineText(text string) string {
	reg := regexp.MustCompile(`end-of-line\d*`)
	return reg.ReplaceAllString(text, "")
}

func TagWithUserName(lines []string) []string {
	reg := regexp.MustCompile(`User\s+(\w*)`)
	var ret []string
	for _, line := range lines {
		submatch := reg.FindStringSubmatch(line)
		if len(submatch) != 0 {
			ret = append(ret, fmt.Sprintf("[USR] %s %s", submatch[1], line))
		} else {
			ret = append(ret, line)
		}
	}
	return ret
}
