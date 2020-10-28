package main

import (
	"fmt"
	"os"
	"net/http"
	// "bufio"
	// "io/ioutil"
	// "io"
	"strings"
)

func main() {
	urls := os.Args[1:]
	for _, url := range urls {
		if !strings.HasPrefix(url, "http://") {
			url = "http://" + url
		}
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "get %s error; %v\n", url, err)
			os.Exit(1)
		}
		// 分段获取
		// content := ""
		// input := bufio.NewScanner(resp.Body)
		// for input.Scan() {
		// 	content += input.Text()
		// }
		// resp.Body.Close()
		// fmt.Println(content)
		
		// 全部读取
		// b, err := ioutil.ReadAll(resp.Body)
		// if err != nil {
		// 	fmt.Fprintf(os.Stderr, "read response error: %v", err)
		// 	os.Exit(1)
		// }
		// resp.Body.Close()
		// fmt.Printf("%s", b)

		// copy
		// _, err = io.Copy(os.Stdout, resp.Body)
		// if err != nil {
		// 	fmt.Fprintf(os.Stderr, "copy error: %v", err)
		// 	os.Exit(1)
		// }
		// resp.Body.Close()

		fmt.Println(resp.Status)
	}
}