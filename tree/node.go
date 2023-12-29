package tree

import "fmt"

type Node struct {
	Value       int
	Left, Right *Node
}

func (t *Node) GetTreeNode(value int, left *Node, right *Node) *Node {
	return &Node{value, left, right}
}
func (t *Node) Print() {
	if t == nil {
		fmt.Println("nil print method")
		return
	}
	fmt.Println(t.Value)
}
