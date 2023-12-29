package main

import (
	"fmt"
	"gocourse/extend"
	"gocourse/tree"
)

func main() {
	// >>>>>>>>>>>>>>>>>> 使用组合扩展>>>>>>>>>>>>>>>>>>
	node := tree.Node{Value: 10}
	node.Left = &tree.Node{}
	node.Right = &tree.Node{Value: 11}
	node.Right.Left = new(tree.Node) //无论是指针还是结构体都是点语法
	exNode := extend.ExNode{Node: &node}
	exNode.PostTraverse()
	println("")

	//>>>>>>>>>>>>>>>>>> 使用别名扩展>>>>>>>>>>>>>>>>>>
	queue := extend.Queue{10, 9, 8}
	for i := 0; i < 5; i++ {
		queue.Push(i * 10)
	}
	for i := 0; i < 10; i++ {
		pop, r := queue.Pop()
		if r == nil {
			fmt.Println(pop)
		} else {
			r.RuntimeError()
			break
		}
	}

	//>>>>>>>>>>>>>>>>>> 使用内嵌扩展>>>>>>>>>>>>>>>>>>
	node2 := extend.ExNode2{Node: &tree.Node{Value: 10}}
	node2.Left = &tree.Node{}
	node2.Right = &tree.Node{Value: 11}
	node2.Right.Left = new(tree.Node)
	node2.PostTraverse()
	println("")
}
