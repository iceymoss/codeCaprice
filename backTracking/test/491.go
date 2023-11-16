package test

//满足长度:>=2
//递增
//结果集中不能有重复项

func findSubsequences(nums []int) [][]int {
	ans := make([][]int, 0)
	path := make([]int, 0)
	var backTracking func(int)
	backTracking = func(index int) {
		if len(path) >= 2 && index <= len(nums) {
			ans = append(ans, append([]int{}, path...))
		}
		if index > len(nums) {
			return
		}
		used := make(map[int]bool)
		for i := 0; i < len(nums); i++ {
			if used[nums[i]] {
				continue
			}
			if len(path) == 0 || nums[i] >= path[len(path)-1] {
				used[nums[i]] = true
				path = append(path, nums[i])
				backTracking(i + 1)
				path = path[:len(path)-1]
			}
		}
	}
	backTracking(0)
	return ans
}
