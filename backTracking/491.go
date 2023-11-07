package backTracking

func findSubsequences(nums []int) [][]int {
	ans := make([][]int, 0)
	path := make([]int, 0)
	var backtracking func(index int)
	backtracking = func(index int) {
		if index <= len(nums) && len(path) >= 2 {
			ans = append(ans, append([]int{}, path...))
		}

		if index > len(nums) {
			return
		}

		used := make(map[int]bool, len(nums)) // 初始化used字典，用以对同层元素去重
		for i := index; i < len(nums); i++ {
			if used[nums[i]] {
				continue
			}
			if len(path) == 0 || path[len(path)-1] <= nums[i] {
				path = append(path, nums[i])
				used[nums[i]] = true
				backtracking(i + 1)
				path = path[:len(path)-1]
			}
		}
	}
	backtracking(0)
	return ans
}
