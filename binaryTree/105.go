package binaryTree

func perBuildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}
	var build func([]int, []int) *TreeNode
	build = func(in, per []int) *TreeNode {
		if len(in) == 0 || len(per) == 0 {
			return nil
		}

		tagPer := per[0]
		per = per[1:]
		node := &TreeNode{Val: tagPer}

		left := getIndexForInorder(tagPer, in)

		node.Left = build(in[:left], per[:left])
		node.Right = build(in[left+1:], per[left:])
		return node
	}
	return build(inorder, preorder)
}

func getIndexForInorder(tag int, arr []int) int {
	for i, v := range arr {
		if v == tag {
			return i
		}
	}
	return -1
}
