# 1.整型
```go
f := 3.141
i := int(f)
fmt.Println(f, i) // "3.141 3"
```

# 2. 浮点数
`float32` 的有效bit位只有23个，其他的bit位用于指数和符号，当整数大于23Bit能表达的范围时，`float32`的表示会出现误差
```go
var f float32 = 16777216 // 1 << 24
fmt.Println(f == f+1) // "true"
```

# 3. 复数
```go
var x complex128 = complex(1, 2) // 1 + 2i
var y complex128 = complex(3, 4) // 3 + 4i
fmt.Println(x*y)            //  -5+10i
```

# 4. 布尔型
`&&` (AND) 和 `||` (OR) 操作符结合，并且有短路行为。  
布尔值不会隐式的转换为数字值0或1，需要自己写一个if来判断返回 `btoi`
```go
func btoi(b bool) int {
    if b {
        return 1
    }
    return 0
}

func itob(i int) bool{
    return i != 0
}
``` 

# 5. 字符串
```go
s := "left foot"
t := s
s += ", right foot"
fmt.Println(s) // "left foot, right foot"
fmt.Println(t) // "left foot"
```
`字符串是不可修改的`
```go
s[0] = 'L' // compile error: cannot assign to s[0]
```

## 5.1 Unicode
Unicode码点对应Go语言中的rune整数类型(int32)

## 5.2 UTF-8
UTF-8 是将 Unicode码点编码为字节序列的变长编码

## 5.3 字符串和Byte切片
实现`basename` 系统路径的文件名，删除前缀和后缀
```go
fmt.Println(basename("a/b/c.go"))   // "c"
fmt.Println(basename("c.d.go"))     // "c.d"
fmt.Println(basename("abc"))        // "abc"

func basename(s string) string{
    slash := strings.LastIndex(s, "/") // -1 if "/" not found
    s = s[slash + 1]
    if dot := strings.LastIndex(s, "."); dot >= 0 {
        s = s[:dot]
    }
    return s
}
```
字符串和字节slice之间可以相互转换
```go
s := "abc"
b := []byte(s)
s2 := string(b)
```
`strings` 包中的六个函数
```go
func Contains(s, substr string) bool
func Count(s, sep string) int
func Fields(s string) []string
func HasPrefix(s, prefix string) bool
func Index(s, sep string) int
func Join(a []string, sep string) string
```
`bytes` 包中的六个函数
```go
func Contains(b, subsline []byte) bool
func Count(s, sep []byte) int
func Fields(s []byte) []string
func HasPrefix(s, prefix []byte) bool
func Index(s, sep []byte) int
func Join(s [][]byte, sep []byte) []byte
```
`bytes.Buffer` 提供用于字节slice的缓存
```go
var buf = bytes.Buffer
buf.WriteByte('a')
buf.WriteString(", ")
buf.String()
```
# 5.4 字符串和数字的转换
```go
x := 123
y := fmt.Sprintf("%d", x) // 返回一个格式化的字符串
y := strconv.Itoa(x)        // 整数到ASCII
```
`FormatInt` 和 `FormatUint` 可以使用不同进制转化数字
```go
strconv.FormatInt(int64(x), 2) // "1111011"

fmt.Sprintf("x=%b", x)          //"x=1111011"
```
字符串到整数
```go  
x, err := strconv.Atoi("123")   // x is an int
y, err := strconv.ParseInt("123", 10, 64) // base 10, up to 64 bits
```

# 6. 常量
```go
const pi = 3.14159
```
一个常量的声明也可以包含一个类型和一个值
```go
const noDelay time.Duration = 0
const timeout = 5 * time.Minute
fmt.Println("%T %[1]v\n", noDelay)          // "time.Duration 0"
fmt.Println("%T %[1]v\n", timeout)           // "time.Duration 5m0s"
fmt.Println("%T %[1]v\n", time.Minute)       // "time.Duration 1m0s"
```
批量声明的常量，除第一个外其他的都可以省略
```go
const (
    a = 1
    b
    c = 2
    d
)
fmt.Println(a, b, c, d) // "1 1 2 2"
```
## 6.1 `iota` 常量生成器
第一个声明的常量所在的行，iota会将其置为0，其他的行加一
```go
const (
    Sunday int = itoa   // 0
    Monday              // 1
    Tuesday             // 2
)

1 << itoa // 第一行 1*2^0, 第二行 1*2^1 ...
```


