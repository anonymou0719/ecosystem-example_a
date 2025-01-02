// project-a/main.go
package main

import (
	"fmt"
	"project-b/example" // 引用项目 B 的 example 包
)

func main() {
	// 调用项目 B 中的 FunctionB 方法
	message := example.FunctionB()
	fmt.Println(message)
}
