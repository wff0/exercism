package twelvedays

import (
	"fmt"
	"strings"
)

var m = map[int]string{
	1:  "first",
	2:  "second",
	3:  "third",
	4:  "fourth",
	5:  "fifth",
	6:  "sixth",
	7:  "seventh",
	8:  "eighth",
	9:  "ninth",
	10: "tenth",
	11: "eleventh",
	12: "twelfth",
}

var animals = []string{
	"a Partridge",
	"two Turtle Doves",
	"three French Hens",
	"four Calling Birds",
	"five Gold Rings",
	"six Geese-a-Laying",
	"seven Swans-a-Swimming",
	"eight Maids-a-Milking",
	"nine Ladies Dancing",
	"ten Lords-a-Leaping",
	"eleven Pipers Piping",
	"twelve Drummers Drumming",
}

func Verse(i int) string {
	sb := new(strings.Builder)
	sb.WriteString(fmt.Sprintf("On the %s day of Christmas my true love gave to me: ", m[i]))
	if i == 1 {
		sb.WriteString("a Partridge in a Pear Tree.")
		return sb.String()
	}
	for start := i; start > 1; start-- {
		sb.WriteString(animals[start-1])
		sb.WriteString(", ")
	}
	sb.WriteString("and a Partridge in a Pear Tree.")
	return sb.String()
}

func Song() string {
	verses := make([]string, 0, 12)
	for i := 1; i <= 12; i++ {
		verses = append(verses, Verse(i))
	}
	return strings.Join(verses, "\n")
}
