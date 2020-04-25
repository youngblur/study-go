# 1. 函数的声明
包括: 函数名、 形式参数列表、返回值列表(可省略) 以及 函数体
```go
func name(parameter-list) (result-list){
    body
}

func add(x int, y int) int {return x + y}
func sub(x, y int) (z int) {z = x - y; return}
func first(x int, _ int) int {return x}
```
`_` 强调某个参数未被使用  
没有函数体的函数声明，表示该函数不是以Go来实现的，来定义函数标识符

# 2. 多返回值
```go
links, err := findLinks(url)

# 可以先打印出log信息， 
func findLinksLog(url string) ([]string, error){
    log.Printf("findLinks %s", url)
    return findLinks(url)
}
```
如果一个函数将返回值都显示了变量名，name该函数return语句可以省略,如 `sub()`

#3. 错误
对于运行失败的函数，它们都会返回一个额外的返回值来传递错误信息。返回值是 `error` (接口)类型。  
error 类型可能是 `nil` (运行成功) 和 `non-nil` (运行失败)。  
这些错误的信息被认为是一种预期的值而非异常.

## 3.1 错误处理策略
- 传播错误。 子程序的失败直接return 变成函数的失败
    ```go
    doc, err := html.Get(url)
    if err != nil {
      return nil, err
    }       
    ```
- 在时间范围内重试。
    ```go
    func WaitForServer(url string) error {
        const timeout = 1 * time.Minute
        deadline := time.Now().Add(timeout)
        for tries := 0; time.Now().Before(deadline); tries++ {
            _, err = := http.Head(url)
            if er == nil {
                return nil
            }
            log.Printf("server not responding (%s); retrying...", err)
            time.Sleep(time.Second << unit(tries))
        }  
        return fmt.Error("server %s failed to respond after %s", url, timeout) 
    }
    ```

- 输出错误，结束程序
    ```go
    if err := WaitForServer(url); err != nil {
        fmt.Fprintf(os.Stderr, "Site is down: %v\n", err)
        os.Exit(1)
    }
    ```
  
- 输出错误，但不终止程序
- 忽略错误

## 3.2 文件结尾错误 (EOF)
`io.EOF`， 任何由文件结束引起的读取失败都返回同一个错误

# 4. 函数值
在go中， 函数被看作为第一类值； 拥有类型，可以被赋值给其他变量，传递给函数，从函数返回。
```go
func square(n int) int {return n * n}

f := square
fmt.Println(f(3)) // "9"

var a func(int) int
a(3)             // 此处 a 为 nil， 会报错
```

# 5. 匿名函数
没有函数名
```go
// square 返回一个匿名函数，这个匿名函数的返回值是int类型
func square() func() int {
    var x int
    return func() int {
        x++
        return  x * x
    }
}

f := square()
fmt.Printf(f()) //  "1"
fmt.Printf(f()) //  "4"
```

## 5.1 警告: 捕获迭代变量
在for循环体内，使用匿名函数是共享 循环变量的，所以当循环结束后，在对匿名函数的list进行操作，保存匿名函数结果都会引用循环最后的一次循环变量。可以在循环中拷贝一个副本 `v := v`.

# 6. 可变参数
在参数列表的最后一个参数类型之前加上省略符号 `...` ，表示该函数会接受任意数量该类型参数
```go
func sum(vals...int) int {
    total := 0
    for _, val := range vals{
        total += val
    }
    return total
}

// vals 被看作是 []int 的切片
fmt.Println(sum())   // "0"
fmt.Println(sum(1,2,3,4)) // "10"
```
虽然类型看上去像切片，但是是不同的
```go
func f(...int) {}
func g([]int) {}
fmt.Printf("%T\n", f) // "func(...int)"
fmt.Printf("%T\n", g) // "func([]int)"
```

# 7. Deferred 函数
在普通函数或者方法前加上关键字 `defer`。 跟在后面的语句会被延迟执行。直到包含该`defer`语句的函数执行完毕后，defer后的语句才会执行。
```go
func double(x int) (result int){
    defer func() {fmt.Printf("double(%d) = %d\n", x, result) }()
    return x + x
}

_ = double(4)           // "double(4) = 8"

func triple(x) (result int){
    defer func() {result += x}()
    return double(x)
}
fmt.Prinln(triple(4))    // "12"
```
使用循环体时需要注意，defer是函数结束后才会执行，所以循环结束并不会立即执行循环体内的defer内容

# 8. Panic 异常
运行时的错误会引起panic异常(数组访问越界、空指针引用等)  
当panic异常发生，会立即执行该goroutime(可以先理解为线程)中被延迟的函数

# 9. Recover 捕获异常
通常来说，不应该对panic异常做任何处理，也许我们可以从异常中恢复，但至少可以在程序崩溃之前做一些事情(如：web服务崩溃前关闭所有连接，不做处理客户端会一直等待)。  
recover 会从异常panic异常中恢复，并返回panic value。在未发生异常时调用recover会返回`nil`
```go 
func Parse(input string) (s *Syntax, err error) {
    defer func() {
        if p := recover(); p != nil {
            err = fmt.Errorf("internal error : %v", p)
        }   
    }()
    // ... parse code ...
}
```


