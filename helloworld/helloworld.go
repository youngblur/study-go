package main  // main 方法才能进行 run

import "fmt"

func main() {
	fmt.Println("Hello, World!")
}

// 使用 go run helloworld.go 来一次性实现
// 使用 go build helloworld.go 会生成可执行文件
