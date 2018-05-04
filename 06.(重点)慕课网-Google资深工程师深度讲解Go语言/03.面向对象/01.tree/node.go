package _1_tree

import (
	"fmt"
)

type TreeNode struct {
	Value       int
	Left, Right *TreeNode
}

// 工厂方法创建对象
func CreateNode(value int) *TreeNode {
	return &TreeNode{Value: value}
}

func (node TreeNode) Print() {
	fmt.Print(node.Value, " ")
}

func (node *TreeNode) SetValue(value int) {
	if node == nil {
		fmt.Println("Setting Value to nil node.Ignored.")
	}
	node.Value = value
}
