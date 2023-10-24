package test

func combine(n int, k int) [][]int {
	ans := make([][]int, 0)
	path := make([]int, 0)
	var backtracking func(index int)
	backtracking = func(index int) {
		if len(path) == k {
			item := append([]int{}, path...)
			ans = append(ans, item)
			return
		}

		if index > n {
			return
		}

		for i := index; i <= n; i++ {
			path = append(path, i)
			backtracking(i + 1)
			path = path[:len(path)-1]
		}
	}
	backtracking(1)
	return ans
}
