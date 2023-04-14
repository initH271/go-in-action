package main

/*
	@name: dup1
	@describe: Receive input from the command line (use ctrl+z+enter to end input), and then output the number of occurrences of each line.
*/
import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// the counts of line
	counts := make(map[string]int)
	// create a input-stream object by taking the std input-stream
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() { // input.Scan() will return a bool "true" if input is detected, else false.
		// record the line
		counts[input.Text()]++
	}
	for line, n := range counts {
		if n > 0 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}
