package binaryTree

type Nodes struct {
	Val   int
	Left  *Nodes
	Right *Nodes
	Next  *Nodes
}

func connect(root *Nodes) *Nodes {
	if root == nil {
		return root
	}
	queue := make([]*Nodes, 0)
	queue = append(queue, root)
	for len(queue) > 0 {
		size := len(queue)
		tmp := make([]*Nodes, 0)
		for size > 0 {
			top := queue[0]
			queue = queue[1:]
			tmp = append(tmp, top)
			if top.Left != nil {
				queue = append(queue, top.Left)
			}
			if top.Right != nil {
				queue = append(queue, top.Right)
			}
			size--
		}
		if len(tmp) > 1 {
			for i := 0; i < len(tmp)-1; i++ {
				tmp[i].Next = tmp[i+1]
			}
		}
	}
	return root
}
