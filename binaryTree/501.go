package binaryTree

func findMode(root *TreeNode) []int {
	tmps := make(map[int]int)
	ans := make([]int, 0)
	var dfs func(*TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		tmps[node.Val]++
		dfs(node.Left)
		dfs(node.Right)
	}
	dfs(root)
	maxCnt := -1
	for _, v := range tmps {
		if v > maxCnt {
			maxCnt = v
		}
	}

	for k, v := range tmps {
		if v == maxCnt {
			ans = append(ans, k)
		}
	}
	return ans
}
