package backTracking

import "sort"

func permuteUnique(nums []int) [][]int {
	ans := make([][]int, 0)
	path := make([]int, 0)
	used := make(map[int]bool)
	sort.Ints(nums)
	var backTracking func()
	backTracking = func() {
		if len(path) == len(nums) {
			ans = append(ans, append([]int{}, path...))
			return
		}
		for i := 0; i < len(nums); i++ {
			if i > 0 && nums[i] == nums[i-1] && !used[i-1] {
				continue
			}
			if !used[i] {
				path = append(path, nums[i])
				used[i] = true
				backTracking()
				path = path[:len(path)-1]
				used[i] = false
			}
		}
	}
	backTracking()
	return ans
}

//
//func deleteDuplicates(arr [][]int, nums []int) bool {
//	for i := 0; i < len(arr); i++ {
//		index := 0
//		for len(arr[i]) == len(nums) && index < len(nums) {
//			if arr[i][index] != nums[index] {
//				return false
//			}
//			index++
//		}
//	}
//	return true
//}
