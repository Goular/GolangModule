package _1_tree

import (
	"fmt"
)

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

func (node *TreeNode) TraverseWithChannel() chan *TreeNode {
	out := make(chan *TreeNode)
	go func() {
		// 这是一个同步顺序递归方法，所以最后才会去执行 close(out)
		node.TraverseFunc(func(node *TreeNode) {
			out <- node
		})
		// 这个仅仅会执行一次，不会执行多次的，不用担心
		close(out)
	}()
	return out
}
