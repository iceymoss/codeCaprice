package main

import (
	"codeCaprice/backTracking"
	"fmt"
)

func main() {
	//s3 := make([]int, 2, 10)
	//fmt.Println(s3) // [0 0]
	//Test2(&s3)
	//fmt.Println(s3) // [0 0]
	//
	//fmt.Println(s3[0:4])
	//s4 := s3[0:5]
	//fmt.Println("s4:", s4)
	//fmt.Println(s4[0:10])
	//fmt.Println("s4:", s4)
	//s4[0] = 100
	//fmt.Println(s3) // [100 0 6 6 6 0 0 0 0 0]
	//fmt.Println(s4) // [100 0 6 6 6 0 0 0 0 0]
	//arr := []int{2, 3, 5}
	//bk.CombinationSum(arr, 8)
	//fmt.Println(strconv.Itoa(-1))
	backTracking.SubsetsWithDup([]int{1, 1, 2, 2})
}
func Test2(s *[]int) {
	*s = append(*s, 6)
	*s = append(*s, 6)
	*s = append(*s, 6)
	fmt.Println(*s) // [0 0 6 6 6]
}
