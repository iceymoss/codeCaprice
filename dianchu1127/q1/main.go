package main

import (
	"fmt"
	"strconv"
)

func compress(str []string) []string {
	ans := make([]string, 0)
	for _, s := range str {
		var res string
		s += " "
		l := 0
		for r := 0; r < len(s); r++ {
			if s[l] == s[r] {
				continue
			}
			count := r - l
			if count > 1 {
				res += string([]byte{s[l]}) + strconv.Itoa(r-l)
			} else {
				res += string([]byte{s[l]})
			}
			l = r
		}
		ans = append(ans, res)
	}
	return ans
}

func main() {
	fmt.Println(compress([]string{"AAABBAADFGDSSSSFFEEFD", "DSAFIHDFIDRFDFFFFFDIJIJIIII", "DIFJDISAKKKKKWWWWWEEEEELLLLLLDUFDNFEISNSSSSSS"}))
}
