package extend

import (
	"cn.chriswood/gocourse/tree"
	"fmt"
)

// ExNode 使用组合来扩展
type ExNode struct {
	Node *tree.Node
}

func (e *ExNode) PostTraverse() {
	if e == nil || e.Node == nil {
		return
	}
	wrapLeft := ExNode{e.Node.Left}
	wrapLeft.PostTraverse()
	wrapRight := ExNode{e.Node.Right}
	wrapRight.PostTraverse()
	fmt.Print(e.Node.Value, " ")
}
