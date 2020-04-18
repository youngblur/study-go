# 声明
- `var` 变量
- `const` 常量
- `type` 类型
- `func` 函数

# 变量
```go
var 变量名 类型 = 表达式

var s string
fmt.Println(s) // ""

var b, f, s = true, 2.3, "four"

var f, err = os.Open(name)
```
## 简短变量声明
在`函数内部`可以简短变量声明语句  
```go
名字 := 表达式

t := 0.0

i, j := 0, 1
i, j = j, i // 交换 i 和 j 的值
````
`:=` 是一个声明语句， 而 `=` 是一个变量赋值操作

## 指针
一个指针的值是另一个变量的地址。一个指针对应变量在内存中的存储位置。并不是每一个值都会有一个内存地址，但是对于每一个变量必然有对应的内存地址。
```go
x := 1
p := &x         // p指向 x 
fmt.Println(*p) // "1"
*p = 2          
fmt.Println(x) // "2"
```
当指针指向同一个变量或全部是nil时才相等
```go
func f *int {
    v := 1
    return &v
}
fmt.Println(f() == f()) // "false", 每次调用返回不同的地址
```
`flag` 包，类似于python的argparse
```go
var sep = flag.String("s", " ", "separator") // 标志名, 默认值， -h 帮助
```

## new 函数
创建变量，返回指针

# 赋值
`=` 更新一个变量的值
## 元组赋值
```go
x, y = y, x+y
a[i], a[j] = a[j], a[i]

v, ok = m[key]
v, ok = <-ch
```
## 可赋值性
```go
mdedals := []string{"gold", "silver", "bronze"}
```

# 类型
```go
type 类型名字 底层类型

type Celsius float64
var c Celsius

// 申明类型 c 的一个方法
func (c Celsius) String() string {return fmt.Sprintf("%g C", c)}
```

# 包和文件
包通过控制哪些名字是外部可见的来隐藏内部实现信息。  
如果一个(变量，常量，类型，函数)名字是大写字母开头的，那么该名字是导出的。  
不同文件可以指定相同包

## 导入包
从 `src` 后的路径开始

## 包的初始化
```go
var a = b + c   // 第三个初始化
var b = c + 1   // 第二个初始化
var c = 1       // 第一个初始化

// 每个文件可以包含多个init初始化函数
func init() {
}
```

## 作用域
```go
func main() {
    x := "hello"
    for _, x := range x {
        x := x + 'A' - 'a'
        fmt.Printf("%c", x) // "HELLO"， 变为大写
    }
}
```



