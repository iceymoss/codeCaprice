package binaryTree

func isValidBST(root *TreeNode) bool {
	ans := make([]int, 0)
	var dfs func(*TreeNode, *[]int)
	dfs = func(node *TreeNode, arr *[]int) {
		if node == nil {
			return
		}

		dfs(node.Left, arr)
		*arr = append(*arr, node.Val)
		dfs(node.Right, arr)
	}
	dfs(root, &ans)
	for i := 0; i < len(ans)-1; i++ {
		if ans[i] >= ans[i+1] {
			return false
		}
	}
	return true
}

func isValidBSTV(root *TreeNode) bool {
	var dfs func(*TreeNode, *[]int) bool
	dfs = func(node *TreeNode, arr *[]int) bool {
		if node == nil {
			return true
		}

		l := dfs(node.Left, arr)
		if len(*arr) != 0 && (*arr)[len(*arr)-1] >= node.Val {
			return false
		}
		*arr = append(*arr, node.Val)
		r := dfs(node.Right, arr)
		return l && r
	}
	return dfs(root, &[]int{})
}
