package extend

import (
	"cn.chriswood/gocourse/tree"
	"fmt"
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
func (t *ExNode2) Traverse() {
	fmt.Println("this is ExNode2's shadowed method Traverse")
}

// 内嵌方式看起来和java、c++的继承相似，实质不一样
// 看起来一样可以重载方法
// 父类引用指向子类对象 在go中内嵌类型扩展无法做到，go中内嵌类型扩展只是组合扩展的语法糖
// go中通过接口来实现类似这种能力
