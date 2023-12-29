package extend

import (
	"fmt"
	"gocourse/tree"
)

// ExNode2 使用内嵌类型来扩展
type ExNode2 struct {
	*tree.Node
}

func (e *ExNode2) PostTraverse() {
	if e == nil || e.Node == nil {
		return
	}
	wrapLeft := ExNode2{e.Left} //注意这里 直接就可以点语法点出内嵌的struct的成分
	wrapLeft.PostTraverse()
	wrapRight := ExNode2{e.Right} //注意这里 直接就可以点语法点出内嵌的struct的成分
	wrapRight.PostTraverse()
	fmt.Print(e.Value, " ") //注意这里 直接就可以点语法点出内嵌的struct的成分
}
