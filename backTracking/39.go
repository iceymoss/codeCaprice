package backTracking

func CombinationSum(candidates []int, target int) [][]int {
	ans := make([][]int, 0)
	var backtracking func(start int, tmp []int, total int)
	backtracking = func(start int, tmp []int, total int) {
		//终止条件是: sum >= target
		if total == target {
			ans = append(ans, append([]int{}, tmp...))
			return
		}

		if total > target {
			return
		}

		//这里需要剪枝：因为可能出现重复，如ans: [[2 2 2 2] [2 3 3] [3 2 3] [3 3 2] [3 5] [5 3]]
		//只要限制下一次选择的起点，是基于本次的选择，这样下一次就不会选到本次选择同层左边的数。
		//基于此，继续选择，传i，下次就不会选到i左边的数
		for i := start; i < len(candidates); i++ {
			tmp = append(tmp, candidates[i])
			backtracking(i, tmp, total+candidates[i])
			//回溯
			tmp = tmp[:len(tmp)-1]
		}
	}
	backtracking(0, []int{}, 0)
	return ans
}
