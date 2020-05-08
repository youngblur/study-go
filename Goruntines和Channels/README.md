#1. Goroutines
- go中，每一个并发的执行单元叫一个 goroutine
- main是一个单独的goroutine，新的goroutine会 在普通函数或方法前加go 来调用
```go
f()   // 调用 f(), 等待返回值
go f() // 创造一个新的 goroutine, 不需要等待
```
`运行程序spinner`,
main 程序调用使用递归的斐波拉契数列，需要一定的时间。期间使用另一个程序打印小图标

#2. 示例：并发的Clock服务
- `clock1`   作为服务端每隔一秒就将当前时间写到客户端
- 可以使用shell自带netcat(nc命令)，来执行网络连接 `nc localhost 8000`

#3. 示例: 并发的Echo服务
略

# 4. Channels



