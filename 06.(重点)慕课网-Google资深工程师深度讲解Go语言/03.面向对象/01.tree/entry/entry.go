package main

import "fmt"
import "baseLearn/03.面向对象/01.tree"

//拓展别人的内容或者是拓展官方包内容
type myTreeNode struct {
	node *_1_tree.TreeNode
}

func (node *myTreeNode) postOrder() {
	if node == nil || node.node == nil {
		return
	}
	left := myTreeNode{node.node.Left}
	left.postOrder()
	right := myTreeNode{node.node.Right}
	right.postOrder()
	node.node.Print()
}

func main() {
	//创建TreeNode
	var root _1_tree.TreeNode // 默认会设定zero值
	fmt.Println(root)

	// 创建根节点
	root = _1_tree.TreeNode{Value: 3}
	// 创建左边节点
	root.Left = &_1_tree.TreeNode{}
	// 创建右边节点
	root.Right = &_1_tree.TreeNode{5, nil, nil}
	// 创建孙子节点
	root.Right.Left = new(_1_tree.TreeNode)
	root.Left.Right = _1_tree.CreateNode(2)

	// 打印当前的tree对象结构
	//nodes := []TreeNode{
	//	{value: 3},
	//	{},
	//	{6, nil, &root},
	//}
	//fmt.Println(nodes)

	//root.print()

	root.Right.Left.SetValue(4)
	root.Right.Left.Print()
	fmt.Println()

	pRoot := &root
	pRoot.Print()
	fmt.Println()
	pRoot.SetValue(500)
	pRoot.Print()

	var pRoot2 *_1_tree.TreeNode
	//pRoot2.setValue(250)
	pRoot2 = &root
	fmt.Println()
	pRoot2.SetValue(560)
	pRoot2.Print()
	fmt.Println()
	root.Traverse()
	fmt.Println()
	myRoot := myTreeNode{&root}
	myRoot.postOrder()
	fmt.Println()
}
