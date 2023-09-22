package binaryTree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func buildTree(inorder []int, postorder []int) *TreeNode {
	if len(postorder) == 0 {
		return nil
	}
	var build func([]int, []int) *TreeNode
	build = func(in []int, post []int) *TreeNode {
		//如果数组为空，则构造完成
		if len(post) == 0 || len(in) == 0 {
			return nil
		}

		//构造node
		tagPost := post[len(post)-1]
		node := &TreeNode{Val: tagPost}
		post = post[:len(post)-1]

		//找到根节点在中序遍历数组的位置
		//从中序遍历中找到一分为二的点，左边为左子树，右边为右子树
		indexLeft := findIndex(tagPost, in)

		//post[:indexLeft],post[indexLeft:]将后续遍历一分为二，左边为左子树，右边为右子树
		node.Left = build(in[:indexLeft], post[:indexLeft])
		node.Right = build(in[indexLeft+1:], post[indexLeft:])
		return node
	}
	return build(inorder, postorder)
}

// 找到根节点在中序遍历数组的位置
func findIndex(tag int, arr []int) int {
	for i, v := range arr {
		if tag == v {
			return i
		}
	}
	return -1
}
