package backTracking

import "sort"

func combinationSum2(candidates []int, target int) [][]int {
	ans := make([][]int, 0)
	arr := make([]int, 0)
	//同层元素值不能重复的，使用used来记录同层,为了方便判断，先排序然然后判断相邻即可
	sort.Ints(candidates)
	used := make([]bool, 0)
	var backtracking func(index int, sum int)
	backtracking = func(index int, sum int) {
		if sum == target {
			ans = append(ans, append([]int{}, arr...))
			return
		}

		if sum > target {
			return
		}

		for i := index; i < len(candidates); i++ {
			// used[i - 1] == true，说明同一树枝candidates[i - 1]使用过
			// used[i - 1] == false，说明同一树层candidates[i - 1]使用过
			if i > 0 && candidates[i] == candidates[i-1] && used[i-1] == false {
				continue
			}
			used[i] = true
			arr = append(arr, candidates[i])
			backtracking(i+1, sum+candidates[i])
			used[i] = false
			arr = arr[:len(arr)-1]
		}
	}
	backtracking(0, 0)
	return ans
}
