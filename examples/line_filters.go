//line filter -  reads input on stdin, processes it, and then prints some derived result to stdout.
//eg - grep and sed

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	//creation of a filter that capitalizes all input text
	scanner := bufio.NewScanner(os.Stdin)

	//traverse tokens
	for scanner.Scan() {
		ucl := strings.ToUpper(scanner.Text())
		fmt.Println(ucl)
	}

	//errors during scan
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}
}
