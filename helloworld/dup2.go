package main

import (
	"bufio"
	"fmt"
	"os"
)

// 读取文件的行数
// Stdin, Stdout, Stderr, 标准输入，输出，错误
func main() {
	counts := make(map[string]int)
	files := os.Args[1:]
	if len(files) == 0{
		countLines(os.Stdin, counts)
	}else{
		for _, file := range files {
			f, err := os.Open(file) // err == nil, 才是文件被正确打开
			if err != nil {
				fmt.Fprint(os.Stderr, err)
				continue
			}
			countLines(f, counts)
			f.Close()
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Println(n, line)
		}
	}
}

func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
	}
}
