package binaryTree

func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 { //数组为空返回空节点
		return nil
	}

	mid := len(nums) / 2
	node := &TreeNode{Val: nums[mid]}

	node.Left = sortedArrayToBST(nums[:mid])
	node.Right = sortedArrayToBST(nums[mid+1:])

	return node
}
