package tree

import "fmt"

// Traverse
// 中序遍历树结构
// Node结构体的方法可以分散在同一个包下的多个文件中
//
//	func demo(){
//	    fmt.Println("随意写的注释中的代码")
//	}
//
// see also GetTreeNode
func (t *Node) Traverse() {
	if t == nil {
		return
	}
	fmt.Print(t.Value, " ")
	t.Left.Traverse()
	t.Right.Traverse()
}
