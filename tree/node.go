package tree

import "fmt"

// Node
// Node结构体定义
type Node struct {
	Value       int
	Left, Right *Node
}

// GetTreeNode
// 获取树的节点
func (t *Node) GetTreeNode(value int, left *Node, right *Node) *Node {
	return &Node{value, left, right}
}

// Print
// 打印当前节点
func (t *Node) Print() {
	if t == nil {
		fmt.Println("nil print method")
		return
	}
	fmt.Println(t.Value)
}
