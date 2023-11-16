package backTracking

func permute(nums []int) [][]int {
	ans := make([][]int, 0)
	path := make([]int, 0)

	//同一条递归分支下，不能有重复的数据
	used := make(map[int]bool)
	var backTracking func()
	backTracking = func() {
		if len(path) == len(nums) {
			ans = append(ans, append([]int{}, path...))
			return
		}
		for i := 0; i < len(nums); i++ {
			if used[nums[i]] {
				continue
			}
			path = append(path, nums[i])
			used[nums[i]] = true
			backTracking()
			used[nums[i]] = false
			path = path[:len(path)-1]
		}
	}
	backTracking()
	return ans
}
