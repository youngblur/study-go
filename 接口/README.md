
# 0. 前言
- 本章是利用 [gophernotes](https://github.com/gopherdata/gophernotes#mac) 在 jupyter上运行的结果
- 利用jupytext将 .ipynb 转化为 .md

# 1. 接口约定
- 接口是一种抽象类型。它不会暴露出所代表的对象的内部值的结构，只会展现它自己的方法


```go
package fmt
import (
    "io"
    "os"
    "bytes"
)

func Fprintf(w io.Writer, format string, args ...interface{}) (int, error) 
func Printf(format string, args ...interface{}) (int, error) {
    return Fprintf(os.Stdout, format, args...) 
}
func Sprintf(format string, args ...interface{}) string { 
    var buf bytes.Buffer
    Fprintf(&buf, format, args...)
    return buf.String()
}
```

# 2. 接口类型


```go
type Reader interface {
    Read(p []byte) (n int, err error)
}
type Closer interface {
    Close() error
}

type ReadWriterCloser interface {
    Reader
    Writer
    Closer
}
```


    go/parser internal error: identifier already declared or resolved


# 3. 实现接口的条件
- 一个类型拥有一个接口所需要的所有方法，那么这个类型就实现了这个接口


```go
// 空接口, 我们可以将任意一个赋值类型给空接口

var any interface{}
any = true
any = map[string]int{"one": 1}
any = "hello"
```

# 4. flag.Value 接口


```go
import (
    "time"
    "flag"
    "fmt"
)
var period = flag.Duration("period", 1*time.Second, "sleep period")

func main(){
    flag.Parse()
    fmt.Printf("Sleep for %v...", *period)
    time.Sleep(*period)
    fmt.Println()
}
main()
```

    Sleep for 1s...


可以使用命令行进行交互，设置 `period` 的值, 如
```
go build sleep

./sleep -period 50ms

./sleep -period "1 day"
```

# 5. 接口值


```go
var w io.Writer
fmt.Printf("%T\n", w) // "<nil>"
w = os.Stdout
fmt.Printf("%T\n", w) // "*os.File"
w = new(bytes.Buffer)
fmt.Printf("%T\n", w) // "*bytes.Buffer"
```

    <nil>
    *os.File
    *bytes.Buffer





    14 <nil>



# 6. sort.Interface 接口
一个内置的排序算法需要知道三个东西：
- 序列的长度
- 表示两个元素比较的结果
- 一种交换两个元素的方式


```go
type Interface interface {
    Len() int
    Less(i, j int) bool
    Swap(i, j int)
}
```


```go
type StringSlice []string
func (p StringSlice) Len() int {return len(p)}
func (p StringSlice) Less(i, j int) bool {return p[i] < p[j]}
func (p StringSlice) Swap(i, j int)  {p[i], p[j] = p[j], p[i]}
```

# 7. http.Handler接口


```go
type Handler interface {
    ServeHTTP(w ResponseWriter, r *Request)
}

func ListenAndServe(address string, h Handler) error
```

# 8. error 接口


```go
type error interface {
    Error() string
}

```


```go
// 创建 error 最简单的方式是调用 errors.New函数

type errorString struct {text string }

func New(text string) error {
    return &errorString{text}
}
func (e *errorString) Error() string { return e.text }
```

# 9. 示例：表达式求值


```go
type Expr interface{} 
```

实现如下示例：
- sqrt(A / pi)
- pow(x, 3) + pow(y, 3)
- (F - 32) * 5 / 9


```go
type Var string // 识别一个变量, e.g.: x

type literal float64 // 识别一个常量， e.g. 3.14

// 一元操作表达式
type unary struct{
    op rune  // one of '+', '-'
    x Expr
}

// 二元操作
type binary struct{
    op rune   // one of '+', '-' ,' *' . '/'
    x, y Expr
}

type call struct{
    fn string   // one of 'pow', 'sin', 'sqrt'
    args []Expr
}
```

需要将变量映射为对应的值


```go
type Env map[Var]float64
```


```go
// 根据变量，返回表达式的值
type Expr interface{
    Eval(env Env) float64
}
```


```go
func (v Var) Eval(env Env) float64{
    return env[v]
}
func (l literal) Eval(_ Env) float64{
    return float64(l)
}
```


```go
func (u unary) Eval(env Env) float64{
    switch u.op {
    case '+':
        return +u.x.Eval(env)
    case '-':
        return -u.x.Eval(env)
    }
    panic(fmt.Sprintf("unsupport unary operator : %q", u.op))
}

func (b binary) Eval(env Env) float64 {
    switch b.op {
    case '+':
        return b.x.Eval(env) + v.y.Eval(env)
    case '-':
        return b.x.Eval(env) - b.y.Eval(env)
    case '*':
        return b.x.Eval(env) * b.y.Eval(env)
    case '/':
        return b.x.Eval(env) / b.y.Eval(env)
    }
    panic(fmt.Sprintf("unsupported binary operator: %q", b.op)) 
}

func (c call) Eval(env Env) float64 { 
    switch c.fn {
    case "pow":
        return math.Pow(c.args[0].Eval(env), c.args[1].Eval(env))
    case "sin":
        return math.Sin(c.args[0].Eval(env))
    case "sqrt":
        return math.Sqrt(c.args[0].Eval(env))
    }
    panic(fmt.Sprintf("unsupported function call: %s", c.fn))
}
```


    repl.go:4:17: invalid qualified type, expecting packagename.identifier, found: u.x.Eval <*ast.SelectorExpr>



```go
import (
    "testing"
    "math"
    "fmt"
)
func TestEval(t *testing.T) { 
    tests := []struct {
        expr string
        env  Env
        want string
    }{
        {"sqrt(A / pi)", Env{"A": 87616, "pi": math.Pi}, "167"}, {"pow(x, 3) + pow(y, 3)", Env{"x": 12, "y": 1}, "1729"}, {"pow(x, 3) + pow(y, 3)", Env{"x": 9, "y": 10}, "1729"}, {"5 / 9 * (F - 32)", Env{"F": -40}, "-40"},
        {"5 / 9 * (F - 32)", Env{"F": 32}, "0"},
        {"5 / 9 * (F - 32)", Env{"F": 212}, "100"},
    }
                             
    var prevExpr string
                             
    for _, test := range tests {
// Print expr only when it changes.
        if test.expr != prevExpr { 
            fmt.Printf("\n%s\n", test.expr) 
            prevExpr = test.expr
        }
        expr, err := Parse(test.expr) 
        if err != nil {
            t.Error(err) // parse error
            continue
        }
        got := fmt.Sprintf("%.6g", expr.Eval(test.env)) 
        fmt.Printf("\t%v => %s\n", test.env, got)
        if got != test.want {
            t.Errorf("%s.Eval() in %v = %q, want %q\n",
            test.expr, test.env, got, test.want)
        }
    }
}
```


    repl.go:25:28: expression returns 1 value, expecting 2: Parse(test.expr)


# 10. 类型断言
- 语法上看起来像 `x.(T)` 被称为断言类型
- 用来检查它操作对象的动态类型是否和断言类型匹配


```go
import "io"

var w io.Writer
w = os.Stdout
f := w.(*os.File)        //success : f == os.Stdout
c := w.(*bytes.Buffer)   // panic: interface holds *os.File
```


    repl.go:6:10: undefined "bytes" in bytes.Buffer <*ast.SelectorExpr>


# 11. 基于类型断言区别错误类型
形如
```
 if some, ok := err.(*someError), ok {
     。。。
 }
```

# 12. 通过类型断言询问行为

# 13. 类型开关


```go
func sqlQuote(x interface{}) string { 
    switch x := x.(type) {
    case nil:
        return "NULL"
    case int, uint:
        return fmt.Sprintf("%d", x) // x has type interface{} here. 
    case bool:
        if x {
            return "TRUE"
        }
        return "FALSE"
    case string:
        return sqlQuoteString(x) // (not shown)
    default:
        panic(fmt.Sprintf("unexpected type %T: %v", x, x)) 
    }
}
```

# 14. 示例：基于标记的 XML解码
略


```go

```
