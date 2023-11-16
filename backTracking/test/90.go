package test

import "sort"

func subsetsWithDup(nums []int) [][]int {
	ans := make([][]int, 0)
	path := make([]int, 0)
	var backtracking func(int)
	sort.Ints(nums)
	backtracking = func(index int) {
		if index <= len(nums) {
			ans = append(ans, append([]int{}, path...))
		}
		if index > len(nums) {
			return
		}
		for i := index; i < len(nums); i++ {
			if i != index && nums[i] == nums[i-1] {
				continue
			}
			path = append(path, nums[i])
			backtracking(i + 1)
			path = path[:len(path)-1]
		}
	}
	backtracking(0)
	return ans
}
