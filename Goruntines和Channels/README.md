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
如果说goroutine是go语言的并发体的话，那么channels则是它们之间的通信机制。
可以让一个goroutine通过它给另一个goroutine发送信息。
```go
ch := make(chan int)  // 一个int类型的chan
```
`<-` 发送和接受都是这个运算符
```go
ch <- x   // 发送
x = <- ch   // 接受
close(ch)  // 关闭
```
可以指定 channel 的容量
```go
ch = make(chan int)
ch = make(chan int, 3)  // 容量是 3
```

## 4.1 不带缓存的Channels
一个基于无缓存Channels的发送操作导致发送者goroutine阻塞，直到另一个goroutine在相同的Channels上执行接受操作，
当成功传输后，两个goroutine继续执行后面的语句。
反之，接受操作先发生，那么接受者阻塞，直到传输成功。  
基于无缓存Channel的发送和接受将导致两个goroutine=做一次同步操作。  
并发是不确定两个时间的先后顺序，谁先谁后都可以。  
在 netcat1 中，如果客户端程序关闭，后台的goroutine可能依旧在工作。
所以我们需要让主goroutine等待后台goroutine完成工作后再推出，可以使用一个channel来同步两个goroutine:
```go
func main() { 
    conn, err := net.Dial("tcp", "localhost:8000")
    if err != nil {
        log.Fatal(err)
    }
    done := make(chan struct{})
    go func() {
        io.Copy(os.Stdout, conn)
        log.Println("done")
        done <- struct{}{}
    }()
    mustCopy(conn, os.Stdin)
    conn.Close()
    <-done
}
```

## 4.2 串联的Channels (Pipeline)
一个channel 的输出作为下一个channel的输入。  
```go
// 一个简单的channel， 第一个计数器，第二个做平方，第三个输出结果

func main() {
	naturals := make(chan int)
    squares := make(chan int)
    
    go func() {
        for x := 0; x < 100; x++ {
            naturals <- x
        }  
        close(naturals) 
    }()
    
    go func() {
        for x := range naturals {
            squares <- x * x
        }
        close(squares)
    }()
    
    for x := range squares{
        fmt.Println(x)
    }
}
```

## 4.3 单方向的Channel
```go
func counter (out chan<- int) {
    for x := 0; x < 100; x++ {
        out <- x
    }
    close(out)
}

func squarer(out chan<- int, in <-chan int) {
    for v := range in {
        out <- v * v
    }
    close(out)
}

func printer(in <-chan int) {
    for v := range in {
        fmt.Println(v)
    }
}

func main() {
    naturals := make(chan int)
    squares := make(chan int)
    go counter(naturals)
    go squarer(squares, naturals)
    printer(squares)
}
```
## 4.4 带缓存的Channels
```go
// 容量为3的字符队列
ch = make(chan string, 3)

// 向channel 发送三个值
ch <- "A"
ch <- "B"
ch <- "C"
// 再发送第四个值的话会发送阻塞

// 接收一个值
fmt.Println(<-ch) // "A"

// 那么 A 出列后有剩余1个容量了

fmt.Println(cap(ch)) // 容量 "3"
fmt.Println(len(ch)) // 内部缓存个数 "2"
````

# 5. 并发的循环
```go
func makeThumbnails3(filename []string) {
    ch := make(chan struct{})
    for _, f := range filenames {
        go func(f string) {
            thumbnail.ImageFile(f)  // ignore errors
            ch <- struct{}{}
        }
    }(f)
    
    for range filenames {
        <-ch
    }
}
```
值得注意的是，f 是作为一个值传入的，而不是共享的。  
如果不指明函数的参数 f， 使用共享的话，不能确保是"当前"的那个f

# 6. 示例： 并发的web爬虫
```go
var tokens = make(chan struct{}, 20) // 用来设置并发的最大量

func crawl(url string) []string {
    fmt.Println(url)
    tokens <- struct{}{}   // 增加一个token
    list, err := links.Extract(url)
    <-tokens            // 释放一个token
    if err != nil {
        log.Print(err)
    }
    return list
}

func main() {
    wroklist := make(chan []string)
    var n int
    
    n++
    go func() { worklist <- os.Args[1:] }()
    
    seen := make(map[string]bool)
    
    for ; n > 0; n-- {
        list := <-worklist
        for _, link := range list {
            if !seen[link] {
                seen[link] = true
                n++
                go func(link string) {
                    worklist <- crawl(link)
                }(link)
            }           
        }
    }    
}
```

# 7. 基于select的多路复用
```go
select {
case <-ch1:
    // ...
case x := <-ch2:
    // ... use x ...
case ch3<- y:
    // ...
default:
    //
}
```
# 8. 示例:并发的字典遍历
略

#9. 并发的退出
```go
var done = make(chan struct{})

func cancelled() bool {
    select {
    case <-done:
        return true
    default:
        return false
    }
}

// 类似下面
selec {
case sema <- struct{}{} // 增加一个token
case <-done:
    return nil
}
defer func() { <-sema}() // 释放一个token
// something...







```





