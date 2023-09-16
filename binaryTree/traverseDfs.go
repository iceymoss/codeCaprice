package binaryTree

import (
	"container/list"
	"fmt"
)

// TreeNode 定义二叉树
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// TODO 递归版
func perTraverse(root *TreeNode) {
	//终止条件
	if root == nil {
		return
	}
	//这一层需要处理的逻辑
	fmt.Println(root.Val)
	perTraverse(root.Left)
	perTraverse(root.Right)
}

func perTraverseAns(root *TreeNode) []int {
	//终止条件
	ans := make([]int, 0)
	var perOrder func(*TreeNode)
	perOrder = func(node *TreeNode) {
		if node == nil {
			return
		}
		//这一层需要处理的逻辑
		ans = append(ans, node.Val)
		perTraverse(node.Left)
		perTraverse(node.Right)
	}
	perOrder(root)
	return ans
}

func midTraverse(root *TreeNode) {
	if root == nil {
		return
	}
	midTraverse(root.Left)
	fmt.Println(root.Val)
	midTraverse(root.Right)
}

func postTraverse(root *TreeNode) {
	if root == nil {
		return
	}
	postTraverse(root.Left)
	postTraverse(root.Right)
	fmt.Println(root.Val)
}

// TODO 迭代版
//前序遍历

func perTraverseIterate(root *TreeNode) []int {
	ans := make([]int, 0)
	if root == nil {
		return ans
	}
	stack := make([]*TreeNode, 0)
	stack = append(stack, root)
	for len(stack) > 0 {
		//每一次弹出栈顶
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		ans = append(ans, node.Val)

		//先进后出，所以右孩子先入栈
		//右孩子入栈
		if node.Right != nil {
			stack = append(stack, node.Right)
		}

		//左孩子节点入栈
		if node.Left != nil {
			stack = append(stack, node.Left)
		}
	}
	return ans
}

// 中序遍历
func midTraversalIterate(root *TreeNode) []int {
	ans := make([]int, 0)
	if root == nil {
		return ans
	}

	st := list.New()
	cur := root

	for cur != nil || st.Len() > 0 {
		if cur != nil {
			st.PushBack(cur)
			cur = cur.Left
		} else {
			cur = st.Remove(st.Back()).(*TreeNode)
			ans = append(ans, cur.Val)
			cur = cur.Right
		}
	}

	return ans
}

// 后序遍历
func postTraverseIterate(root *TreeNode) []int {
	ans := make([]int, 0)
	if root == nil {
		return ans
	}
	stack := make([]*TreeNode, 0)
	stack = append(stack, root)
	for len(stack) > 0 {
		//每一次弹出栈顶
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		ans = append(ans, node.Val)

		//左孩子节点入栈
		if node.Left != nil {
			stack = append(stack, node.Left)
		}
		//右孩子入栈
		if node.Right != nil {
			stack = append(stack, node.Right)
		}
	}
	reverse(ans)
	return ans
}

func reverse(ans []int) {
	l, r := 0, len(ans)-1
	for l < r {
		ans[l], ans[r] = ans[l], ans[r]
		l++
		r--
	}
}
