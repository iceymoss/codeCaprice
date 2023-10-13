package backTracking

func combinationSum3(k int, n int) [][]int {
	ans := make([][]int, 0)
	path := []int{}
	//回溯，找到对应长度为k的数组， 然后计算sum(path) == n
	var backtracking func(tag int)
	backtracking = func(tag int) {
		//终止条件
		if len(path) == k {
			sum := 0
			for _, v := range path {
				sum += v
			}
			if sum == n {
				ans = append(ans, append([]int(nil), path...))
				return
			}
		}

		if tag > 9 {
			return
		}

		for i := tag; i <= 9; i++ {
			path = append(path, i)
			backtracking(i + 1)
			//回溯，撤销
			path = path[:len(path)-1]
		}
	}
	backtracking(1)
	return ans
}
