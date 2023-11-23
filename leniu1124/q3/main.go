package main

import (
	"fmt"
)

func minTime(s int) string {
	t1 := s / 1
	t2 := (s / 10) - 10
	if t1 < t2 {
		return "v"
	} else {
		return "w"
	}

}

func main() {
	s := 0
	n, _ := fmt.Scan(&s)
	fmt.Println(minTime(n))
}
