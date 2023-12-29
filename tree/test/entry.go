package main

import (
	"fmt"
	"gocourse/tree"
)

func main() {
	node := tree.Node{Value: 10}
	node.Left = &tree.Node{}
	node.Right = &tree.Node{Value: 11}
	node.Right.Left = new(tree.Node) //无论是指针还是结构体都是点语法
	node.Print()
	node.Left.Left.Print()
	node.Traverse()

	println("")
	var nodes = []tree.Node{
		{1, nil, nil},
		{},
		{666, nil, nil},
	}
	fmt.Println(nodes)
}
