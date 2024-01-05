package extend

import (
	"fmt"
	"github.com/chriswoodcn/gocourse/tree"
	"testing"
)

func TestExtend1(t *testing.T) {
	// >>>>>>>>>>>>>>>>>> 使用组合扩展>>>>>>>>>>>>>>>>>>
	node := tree.Node{Value: 10}
	node.Left = &tree.Node{}
	node.Right = &tree.Node{Value: 11}
	node.Right.Left = new(tree.Node) //无论是指针还是结构体都是点语法
	exNode := ExNode{Node: &node}
	exNode.PostTraverse()
	println("")
}
func TestExtend2(t *testing.T) {
	//>>>>>>>>>>>>>>>>>> 使用别名扩展>>>>>>>>>>>>>>>>>>
	queue := Queue{10, 9, 8}
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

}
func TestExtend3(t *testing.T) {
	//>>>>>>>>>>>>>>>>>> 使用内嵌扩展>>>>>>>>>>>>>>>>>>
	node2 := ExNode2{Node: &tree.Node{Value: 10}}
	node2.Left = &tree.Node{}
	node2.Right = &tree.Node{Value: 11}
	node2.Right.Left = new(tree.Node)
	node2.Traverse()
	node2.Node.Traverse()
	node2.PostTraverse()
	println("")
}
