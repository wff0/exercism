package bottlesong

import (
	"fmt"
	"strings"
)

var tpls = []string{
	`%s green %s hanging on the wall,`,
	`%s green %s hanging on the wall,`,
	`And if one green bottle should accidentally fall,`,
	`There'll be %s green %s hanging on the wall.`,
}

var digit2Num = map[int]string{
	0:  "no",
	1:  "one",
	2:  "two",
	3:  "three",
	4:  "four",
	5:  "five",
	6:  "six",
	7:  "seven",
	8:  "eight",
	9:  "nine",
	10: "ten",
}

func Recite(startBottles, takeDown int) []string {
	if takeDown > startBottles {
		return nil
	}
	res := make([]string, 0)
	for i := startBottles; i > startBottles-takeDown; i-- {
		if i < startBottles {
			res = append(res, "")
		}
		num := digit2Num[i]
		next := digit2Num[i-1]
		for j := range tpls {
			if j == 0 || j == 1 {
				//num = strings.Title(num)
				num = strings.Title(num)
				if i >= 2 {
					res = append(res, fmt.Sprintf(tpls[j], num, "bottles"))
				} else {
					res = append(res, fmt.Sprintf(tpls[j], num, "bottle"))
				}
			} else if j == 3 {
				if i-1 >= 2 || i-1 == 0 {
					res = append(res, fmt.Sprintf(tpls[j], next, "bottles"))
				} else {
					res = append(res, fmt.Sprintf(tpls[j], next, "bottle"))
				}
			} else {
				res = append(res, tpls[j])
			}
		}
	}
	return res
}
