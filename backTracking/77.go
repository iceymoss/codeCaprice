package backTracking

// 回溯：n = 4, k = 2
// 输出：
// [
// [2,4],
// [3,4],
// [2,3],
// [1,2],
// [1,3],
// [1,4],
// ]
func combine(n int, k int) [][]int {
	ans := make([][]int, 0)
	var backtracking func(int)
	path := []int{}
	backtracking = func(tag int) {
		//递归终止条件
		if len(path) == k {
			//因为这里slice引用类型，所以需要将值存储到新空间中
			ans = append(ans, append([]int(nil), path...))
			return
		}

		//防止tag越界
		if tag > n {
			return
		}

		//处理逻辑
		for i := tag; i <= n; i++ {
			//记录符合条件的组合
			path = append(path, i)
			backtracking(i + 1)
			//回溯
			path = path[:len(path)-1]
		}
	}
	backtracking(1)
	return ans
}
