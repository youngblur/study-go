package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	for _, url := range os.Args[1:] {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintln(os.Stderr, "fetch: ",err)
			os.Exit(1)
		}
		// resp的Body 包括一个可读的服务器响应流
		b, err := ioutil.ReadAll(resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Fprint(os.Stderr, "fetch: reading", url, err)
			os.Exit(1)
		}
		fmt.Println(b)
	}
}
