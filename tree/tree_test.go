package tree

import (
	"fmt"
	"testing"
)

// ExampleNode_GetTreeNode 示例代码GetTreeNode
func ExampleNode_GetTreeNode() {
	//此处可以编写示例代码
	fmt.Println("ExampleNode_GetTreeNode")
}

// ExampleNode_Print 示例代码Print
func ExampleNode_Print() {
	//此处可以编写示例代码
	fmt.Println("ExampleNode_Print")
}

// ExampleNode_Traverse 示例代码Traverse
func ExampleNode_Traverse() {
	//此处可以编写示例代码
	fmt.Println("此处可以编写示例代码")
}

func TestTree(t *testing.T) {
	node := Node{Value: 10}
	node.Left = &Node{}
	node.Right = &Node{Value: 11}
	node.Right.Left = new(Node) //无论是指针还是结构体都是点语法
	node.Print()
	node.Left.Left.Print()
	node.Traverse()

	println("")
	var nodes = []Node{
		{1, nil, nil},
		{},
		{666, nil, nil},
	}
	fmt.Println(nodes)
}
