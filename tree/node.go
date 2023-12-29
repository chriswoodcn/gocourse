package main

import "fmt"

type TreeNode struct {
	value       int
	left, right *TreeNode
}

func (t *TreeNode) getTreeNode(value int, left *TreeNode, right *TreeNode) *TreeNode {
	return &TreeNode{value, left, right}
}
func (t *TreeNode) print() {
	if t == nil {
		fmt.Println("nil print method")
		return
	}
	fmt.Println(t.value)
}
func (t *TreeNode) traverse() {
	if t == nil {
		return
	}
	fmt.Print(t.value, " ")
	t.left.traverse()
	t.right.traverse()
}
func main() {
	node := TreeNode{value: 10}
	node.left = &TreeNode{}
	node.right = &TreeNode{11, nil, nil}
	node.right.left = new(TreeNode) //无论是指针还是结构体都是点语法
	node.print()
	node.left.left.print()
	node.traverse()
	println("")
	var nodes = []TreeNode{
		{1, nil, nil},
		{},
		{666, nil, nil},
	}
	fmt.Println(nodes)
}
