package binaryTree

// 直接遍历即可：dfs或者bsf都可以
func searchBST(root *TreeNode, val int) *TreeNode {
	//dfs即可
	var dfs func(*TreeNode) *TreeNode
	dfs = func(node *TreeNode) *TreeNode {
		if node == nil {
			return nil
		}
		if node.Val == val {
			return node
		}
		var Lnode, Rnode *TreeNode
		Lnode = dfs(node.Left)
		Rnode = dfs(node.Right)
		if Lnode != nil {
			return Lnode
		}
		if Rnode != nil {
			return Rnode
		}
		return nil
	}
	return dfs(root)
}

// 优化：二分搜索树，特点：左子树所有节点都小于父亲节点值，右子树所有节点都大于父亲节点

func searchBSTV1(root *TreeNode, val int) *TreeNode {
	var search func(*TreeNode) *TreeNode
	search = func(node *TreeNode) *TreeNode {
		if node == nil {
			return nil
		}
		if node.Val == val {
			return node
		}
		if node.Val > val {
			return search(node.Left)
		}
		if node.Val < val {
			return search(node.Right)
		}
		return nil
	}
	return search(root)
}
