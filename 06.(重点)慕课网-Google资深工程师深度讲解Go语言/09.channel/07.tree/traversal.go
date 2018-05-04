package _1_tree

import "fmt"

// 遍历tree
func (node *TreeNode) Traverse() {
	node.TraverseFunc(func(node *TreeNode) {
		node.Print()
	})
	fmt.Println()
}

func (node *TreeNode) TraverseFunc(f func(node *TreeNode)) {
	if node == nil {
		return
	}
	node.Left.Traverse()
	f(node)
	node.Right.Traverse()
}
