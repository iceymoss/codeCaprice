package test

import "fmt"

func combinationSum3(k int, n int) [][]int {
	ans := make([][]int, 0)
	path := make([]int, 0)
	var backtracking func(index int)
	backtracking = func(index int) {
		//递归终止条件
		if len(path) == k && sum(n, path) {
			item := append([]int{}, path...)
			ans = append(ans, item)
			return
		}
		if index > 9 {
			return
		}
		for i := index; i <= 9; i++ {
			path = append(path, i)
			backtracking(i + 1)
			path = path[:len(path)-1]
		}
	}
	backtracking(1)
	return ans
}

func sum(n int, num []int) bool {
	s := 0
	for _, v := range num {
		fmt.Println(v)
		s += v
	}
	return s == n
}
