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


