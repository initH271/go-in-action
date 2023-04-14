package main

/*
	添加需求：出现重复的行时，打印文件名称
*/
import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		// 从输入中读取
		countLines(os.Stdin, counts)
	} else {
		// 从文件中读取
		for _, file := range files {
			f, err := os.Open(file)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts)

			f.Close()
		}
	}
	for line, n := range counts {
		if n > 0 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

// counts all lines
func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
		if counts[input.Text()] > 1 {
			fmt.Printf("%s: 重复出现(\t%s\t)\n", f.Name(), input.Text())
		}
	}
}
