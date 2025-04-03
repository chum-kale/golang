// has 2 appraoches - functions and firect
// undo occurs directlty in main only.

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func append_to_string(s string, w string) string {
	res := s + w
	return res
}

func del_last_chars(s string, k int) string {
	res := s[:len(s)-k]
	return res
}

func print_char(s string, k int) {
	fmt.Println(s[k])
}

func main() {
	var operations int
	fmt.Scan(&operations)

	scanner := bufio.NewScanner(os.Stdin)
	text := ""
	stack := []string{}

	for i := 0; i < operations; i++ {
		scanner.Scan()
		input := scanner.Text()
		parts := strings.SplitN(input, " ", 2)

		command := parts[0]

		switch command {
		case "1": // append
			stack = append(stack, text) // Save current state
			text += parts[1]
		case "2": // delete
			stack = append(stack, text) // Save current state
			k, _ := strconv.Atoi(parts[1])
			if k > len(text) {
				text = ""
			} else {
				text = text[:len(text)-k]
			}
		case "3": // print
			k, _ := strconv.Atoi(parts[1])
			fmt.Println(string(text[k-1]))
		case "4": // undo
			if len(stack) > 0 {
				text = stack[len(stack)-1]
				stack = stack[:len(stack)-1]
			}
		default:
			fmt.Println("Invalid command")
		}
	}

}
