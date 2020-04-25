# 1. 数组
数组的长度是数组类型的一个组成部分，在编译阶段确定。所以长度必须是常量，且无法更改
```go
var a [3]int        // 都为0
var q [3]int{1, 2, 3}
var r [3]int{1, 2} // r[2] = 0

q := [...]int{1,2,3}  // [3]int, ...是，数组的大小由初始化来决定

r := [...]int{0:1, 1:2} // [1,2]
r := [...]int{99:-1}  // 0-98都是0， 99是-1
```
在函数调用数组时，函数接受的是变量一个复制的副本。所以无法更改数组的值，我们可以显式的传入指针来修改
```go
func zero(ptr *[32]byte){
    for i := range ptr{
        ptr[i] = 0
    } 
}
```

# 2. Slice
Slice(切片)代表变长的序列。 slice的语法和数组很像，只是没有固定长度而已。其底层确实引用了一个数组对象。  
Slice由三个部分组成: 指针、长度 和 容量。
长度对应slice中元素的个数， 长度不能超过容量，容量一般是从slice的开始位置到底层数据的结尾位置。`len`和`cap`分别返回长度 和 容量。
```go
func reverse(s []int) {
    for i, j := 0, len(s)-1; i < j; i, j = i + 1, j - 1{
        s[i], s[j] = s[j], s[i]
    }
}
reverse(a[:])
```
创建 slice， 可以为nil
```go
var s []int  // len(s) == 0, s == nil

make([]T, len)
make([]T, len, cap)
```

## 2.1 `append` 函数
向slice追加元素
```go
var runes []rune
for _, r := range "hello, 世界" {
    runes = append(runes, r)
}
fmt.Printf("%q\n", runes) // "['h', ...,'界']"
```
第一个版本的 `appendInt`
```go
func appendInt(x []int, y int) []int {
    var z []int
    zlen := len(x) + 1
    if zlen <= cap(x) {
        z = x[:zlen]
    }else{
        zcap := zlen
        if zcap < 2 * len(x) {
            zcap = 2 * len(x)
        }   
        z = make([]int, zlen, zcap)
        copy(z, x)
    }
    z[len(x)] = y
    return z
}
```
第二个版本，接受多个y传入
```go
func appendInt(x []int, y ...int) []int {
    var z []int
    zlen := len(x) + len(y)
    copy(z[len(x):], y)
    return z
}
```

## 2.2 Slice内存技巧
判断 非空字符串，可以共享内存
```go
func nonempty(strings []string) []string{
    out := strings[:0] // zero-length slice of original
    for _, s := range strings{
        if s != "" {
            out = append(out, s)
        }
    }
    return out
}

```
移除指定位置的元素
```go
func remove(slice []int, i int) []int{
    copy(slice[i:], slice[i+1:])
    return slice[:len(slice)-1]
}
```
# 3. Map
创建map一个空map
```go
ages := make(map[string]int)
ages := map[string]int{}
```
内置的`delete`函数可以删除元素
```go
delete(ages,"alice")
```
当查找不存在时，将返回零值,也可以使用 `age, ok = ages[key]` 
map中的元素并不是一个变量，因此我们不能对map元素进行取址操作

# 4. 结构体
结构体内部成员顺序不同，则类型不同.可以直接对成员进行赋值。
```go
type Employee struct {
    ID      int
    Name    String
}
var dilbert Employee
// 可以指向结构体，也可以指向结构体的成员
var index *Employee = &disbert
name := &Employee.Name
```
相邻成员拥有相同类型，可以合在一行  
对于`首字母小写`的成员，其是隐式的，未导出的。所以不能导出。  
如果结构的成员都是可以比较的，那么结构体也是可以比较的
## 4.1 结构体嵌入和匿名成员
```go
type Point struct{
    X, Y  int
}

type Circle struct{
    Center Point
    Radius int
}

var c Circle
c.Center.X = 1
c.Center.Y = 2
```
结构体重只声明成员的类型，而不指名成员的名字，这类成员就是匿名成员。
```go
type Circle struct{
    Point
    Radius int
}

var c Circle
c.X = 1     // 省去 Center
c.Y = 1
```
但是不能进行字面值语法，即
```go
c = Circle{1,1,1}  // 不能成功编译

c = Circle{Point{1,1}, 1} 
```

#5. JSON
JSON对象可以用来编码Go语言中的map类型(key 是字符串) 和 结构体
[josn完整例子](json.go)
```go
type Movie = struct{
    Title string
    // 因为go里成员都是大写的，但是有些json是小写的，所以用tag来控制
    Year int `json: "released"`		//Tag 来控制编码和解码，将 Year 对应到 json 中的 released.
    Color bool `json:"color, omitempty"`
    Actors []string
}

func main() {
    var movies = []Movie{
        {Title: "Casablance", Year: 1942, Color: false,
            Actors: []string{"Humphrey Bogart", "Ingrid Bergman"}},
        {Title: "Cool Hand Luke", Year: 1967, Color: true,
            Actors: []string{"Paul Newman"}},
        // ...
    }
    data, err := json.Marshal(movies)
    if err != nil{
        log.Fatal("JSON marshling failed: %s", err)
    }
    fmt.Printf("%s \n", data)
    
    // 便于阅读，prefix 是每一行输出的前缀，indent是每一层输出的前缀
    data2, err2 := json.MarshalIndent(movies, "", "    ")
    if err2 != nil {
        log.Fatal("JSON marshling failed: %s", err)
    }
    fmt.Printf("%s \n", data2)
    
    var titles []struct{ Title string}
    if err3 := json.Unmarshal(data, &titles); err3 != nil {
        log.Fatal("JSON unmarshaling failed: %s", err3)
    }
    fmt.Println(titles)
}
```

# 6. 文本 和 HTML 模板
...








