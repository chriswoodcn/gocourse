package private

import (
	testPrivateRepo "codeup.aliyun.com/6142b6b5ebe08992f3507f10/goRepo/test-private-repo.git"
	"fmt"
	"testing"
)

// 私有库 go mod init codeup.aliyun.com/6142b6b5ebe08992f3507f10/goRepo/test-private-repo.git
// go.mod 显示 module codeup.aliyun.com/6142b6b5ebe08992f3507f10/goRepo/test-private-repo.git
// 仓库地址https://codeup.aliyun.com/6142b6b5ebe08992f3507f10/goRepo/test-private-repo.git

//git@codeup.aliyun.com:6142b6b5ebe08992f3507f10/goRepo/test-private-repo.git
//https://codeup.aliyun.com/6142b6b5ebe08992f3507f10/goRepo/test-private-repo.git

// 配置GoMod私有仓库 go env -w GOPRIVATE="codeup.aliyun.com"
// 配置不加密访问 go env -w GOINSECURE="codeup.aliyun.com"
// 配置不使用代理 go env -w GONOPROXY="codeup.aliyun.com"
// 配置不验证包 go env -w GONOSUMDB="codeup.aliyun.com"
// go get  codeup.aliyun.com/6142b6b5ebe08992f3507f10/goRepo/test-private-repo.git
func TestPrivateUse(t *testing.T) {
	fmt.Println(testPrivateRepo.MyName())
}
