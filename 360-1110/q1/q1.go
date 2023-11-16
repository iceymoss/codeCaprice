package main

import (
	"fmt"
)

//
//func minOperations(nums []int) int {
//	n := len(nums)
//	tmp := find(nums)
//	sort.Ints(nums)
//	sum := make([]int, n+1)
//	for i, x := range nums {
//		sum[i+1] = sum[i] + x
//	}
//	ans := make([]int64, len())
//}
//
//func find(nums []int) int {
//	count := make(map[int]int)
//	for _, num := range nums {
//		count[num]++
//	}
//	mode := -1
//	maxCount := 0
//	for num, c := range count {
//		if c > maxCount {
//			mode = num
//			maxCount = c
//		}
//	}
//	return mode
//}

func main() {
	var n int
	nums := make([]int, 0, n)
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		var num int
		fmt.Scan(&num)
		nums = append(nums, num)
	}
	//fmt.Println(minOperations(nums))
}
