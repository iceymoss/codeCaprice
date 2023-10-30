package backTracking

func subsets(nums []int) [][]int {
	ans := make([][]int, 0)
	path := make([]int, 0)
	var backtracking func(start int)
	backtracking = func(start int) {
		//只要start没越界，都是nums的子集
		if start <= len(nums) {
			ans = append(ans, append([]int{}, path...))
		}

		if start > len(nums) {
			return
		}

		for i := start; i < len(nums); i++ {
			path = append(path, nums[i])
			backtracking(i + 1)
			path = path[:len(path)-1]
		}
	}
	backtracking(0)
	return ans
}
