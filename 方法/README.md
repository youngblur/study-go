
# 0. 前言
- 本文本是利用 [gophernotes](https://github.com/gopherdata/gophernotes#mac) 在 jupyter上运行的结果
- 利用jupytext将 .ipynb 转化为 .md

# 1. 方法申明


```go
import (
    "math"
    "fmt"
)

type Point struct {X, Y float64}

// trainditional function
func Distance(p, q Point) float64 {
    return math.Hypot(q.X - p.X, q.Y - p.Y)
}

// same thing, but as a method of the Point type
func (p Point) Distance(q Point) float64 { 
    return math.Hypot(q.X-p.X, q.Y-p.Y)
}
```

p 叫做方法的接收器， go 没有其他语言那样勇 this 或者 self


```go
p := Point{1, 2}
q := Point{4, 6}
fmt.Println(Distance(p, q))
fmt.Println(p.Distance(q))
```

    5
    5





    2 <nil>



go 语言中能够给`任意类型`增加方法


```go
// 声明 Path 是一个 slice 类型，但是我们也可以为其增加方法
type Path []Point 

func (path Path) Distance() float64 {
    sum := 0.0
    for i, _ := range path {
        if i > 0 {
        sum += path[i-1].Distance(path[i])
        } 
    }
    return sum 
}
```


```go
perim := Path{
    {1, 1},
    {5, 1},
    {5, 4},
    {1, 1},
}
fmt.Println(perim.Distance())
```

    12





    3 <nil>



# 2. 基于指针对象的方法

在更新接收器的值时，可以声明指针的形式


```go
func (p *Point) ScaleBy(factor float64) {
    p.X *= factor
    p.Y *= factor
}

r := &Point{1, 2}
r.ScaleBy(2)
fmt.Println(*r) 
```

    {2 4}





    6 <nil>



在go语言中，若接收器 是一个类型的变量，并且其方法需要这个类型的指针作为接收器，我们可以直接传入这个变量，编译器会自动帮我们去变为其指针


```go
p := Point{1,2}
p.ScaleBy(2)
fmt.Println(p)
```

    {2 4}





    6 <nil>



但是我们不能通过一个无法取得地址的接收器来调用指针，比如使用临时变量的内存无法获取


```go
Point{1, 2}.ScaleBy(2) // compile error
```

`nil` 也是一个合法的接收器类型, 所以注意条件判断

# 3. 通过嵌入结构体来拓展类型


```go
import "image/color"

type ColoredPoint struct {
    Point
    Color color.RGBA
}
```


```go
var cp ColoredPoint
cp.X = 1
fmt.Println(cp.Point.X)  // "1"
cp.Point.Y = 2
fmt.Println(cp.Y)     // “2”
```

    1
    2





    2 <nil>



ColoredPoint 同样也有 Point类的方法，可以直接调用


```go
// 简单的cache
import "sync"

var (
    mu sync.Mutex    // mutex 互斥量
    mapping = make(map[string]string)
)

func Lookup(key string) string {
    mu.Lock()
    v := mapping[key]
    mu.Unlock()
    return v
}

// 可以重写为
var cache = struct {
    sync.Mutex
    mapping map[string]string
}{
    mapping: make(map[string]string),
}

func Lookup(key string) string {
    cache.Lock()
    v := cache.mapping[key]
    cache.Unlock()
    return v
}
```


    repl.go:25:5: undefined "cache" in cache.Lock <*ast.SelectorExpr>


# 4. 方法值和方法表达式


```go
p := Point{1, 2}
q := Point{4, 6}
distanceFromP := p.Distance
fmt.Println(distanceFromP(q))  // 方法可以直接当值传递
```

    5





    2 <nil>



# 5. 示例：Bit数组

一个bit。数组通常为一个无符号数的slice来表示，每一个元素的每一位都表示集合里的一个值


```go
// An IntSet is a set of small non-negative integers. // Its zero value represents the empty set.
type IntSet struct {
    words []uint64
}
// Has reports whether the set contains the non-negative value x.
func (s *IntSet) Has(x int) bool {
    word, bit := x/64, uint(x%64)
    return word < len(s.words) && s.words[word]&(1<<bit) != 0
}
// Add adds the non-negative value x to the set.
func (s *IntSet) Add(x int) { 
    word, bit := x/64, uint(x%64) 
    for word >= len(s.words) {
        s.words = append(s.words, 0) 
    }
    s.words[word] |= 1 << bit 
}
// UnionWith sets s to the union of s and t.
func (s *IntSet) UnionWith(t *IntSet) { 
    for i, tword := range t.words {
        if i < len(s.words) { 
            s.words[i] |= tword
        } else {
            s.words = append(s.words, tword)
        } 
    }
}
```


    repl.go:8:53: mismatched types in binary operation & between <uint64> and <int>: s.words[word] & (1 << bit)


# 6. 封装
一个对象的变量或者方法对调用方不可见的话，一般定义为"封装"
- 调用方不能直接修改对象的变量值，
- 隐藏实现的细节
- 阻止外部调用方对对象内部的值任意进行修改

struct 结构小些变量名


```go

```
