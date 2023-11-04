package main

import (
	"fmt"
)

func getNoeSum(nums []int, tag int) int {
	var ans int
	for j := 0; j < len(nums); j++ {
		if sum(nums, j, nums[j]) >= tag {
			ans = ans + 1
		}
	}
	return ans
}

func sum(nums []int, index int, tmp int) int {
	s := 0
	for i := 0; i < len(nums); i++ {
		if tmp == nums[i] && i == index {
			continue
		}
		s += nums[i]
	}
	fmt.Println(s)
	return s
}

func main() {
	var num, tag int
	var ans int

	fmt.Scan(&num, &tag)
	fmt.Println(num, tag)
	var nums []int

	for i := 0; i < num; i++ {
		var n int
		fmt.Scan(&n)
		nums = append(nums, n)
	}
	fmt.Println(nums)

	for j := 0; j < len(nums); j++ {
		if sum(nums, j, nums[j]) >= tag {
			ans = ans + 1
		}
	}
	fmt.Println(ans)

	//fmt.Println(getNoeSum([]int{3, 4, 3}, 7))
}
