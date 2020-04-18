// 查找重复的行
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	inputs := bufio.NewScanner(os.Stdin)
	for inputs.Scan() {
		if inputs.Text() == "end" {
			break
		}
		counts[inputs.Text()]++
	}

	for line, n := range counts {
		if n > 1{
			fmt.Println(n, line)
		}
	}
}