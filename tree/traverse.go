package tree

import "fmt"

// Traverse Node结构体的方法可以分散在同一个包下的多个文件中
func (t *Node) Traverse() {
	if t == nil {
		return
	}
	fmt.Print(t.Value, " ")
	t.Left.Traverse()
	t.Right.Traverse()
}
