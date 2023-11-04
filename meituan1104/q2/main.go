package main

import (
	"fmt"
	"math"
)

// 回溯-枚举所有
func backtracking(nums []int) int {
	max := math.MinInt
	ans := make([]int, 0)
	path := make([]int, 0)
	var dfs func(start int)
	dfs = func(start int) {
		if len(path) == 2 {
			tmp := Multiply(nums, path[0], path[1])
			if max < Multiply(nums, path[0], path[1]) {
				max = tmp
			}
			ans = append(ans, tmp)
			return
		}
		if len(path) > 2 {
			return
		}

		for i := start; i < len(nums); i++ {
			path = append(path, i)
			dfs(i + 1)
			path = path[:len(path)-1]
		}
	}
	dfs(0)
	fmt.Println("ans:", ans)
	return max
}

func main() {
	var n int
	fmt.Scan(&n)
	nums := make([]int, n)
	for i := 0; i < n; i++ {
		var num int
		fmt.Scan(&num)
		nums[i] = num
	}

	fmt.Println(nums)

	//暴力枚举
	max := math.MinInt
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			tmp := Multiply(nums, i, j)
			if tmp > max {
				max = tmp
			}
		}
	}
	fmt.Println(max)

	fmt.Println(backtracking(nums))

}

func Multiply(nums []int, i, j int) int {
	s := 1
	for index, v := range nums {
		tmp := v
		if index == i {
			tmp = v + 1
		}
		if index == j {
			tmp = v - 1
		}
		s *= tmp
	}
	return s
}
