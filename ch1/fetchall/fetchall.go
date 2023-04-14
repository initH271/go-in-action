package main

import (
	"fmt"
	"io"
	// "io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"
)

func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err) // 发送到channel
		return                // 结束当前的goroutine
	}
	websiteFile, err := os.Create("last.html")
	if err != nil {
		fmt.Fprintf(os.Stderr, "fetchall: %v\n", err)
		return
	}
	// nBytes, err := io.Copy(ioutil.Discard, resp.Body) // 响应内容丢进垃圾桶，但是需要读取其字节数
	nBytes, err := io.Copy(websiteFile, resp.Body) // 响应内容丢进单独的文件，但是需要读取其字节数

	resp.Body.Close()
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err) // 发送到channel
		return                                              // to end the goroutine
	}

	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %7d bytes %s", secs, nBytes, url)

}

func main() {
	start := time.Now()
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		if !strings.HasPrefix(url, "http://") {
			url = "https://" + url
		}
		go fetch(url, ch) // start a goroutine
	}

	for range os.Args[1:] {
		fmt.Println(<-ch) // 从channel读取消息并打印
	}
	fmt.Printf("耗时: %.2fs \n", time.Since(start).Seconds())
}
