package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

/*
 * Complete the 'highestValuePalindrome' function below.
 *
 * The function is expected to return a STRING.
 * The function accepts following parameters:
 *  1. STRING s
 *  2. INTEGER n
 *  3. INTEGER k
 */

func highestValuePalindrome(s string, n int32, k int32) string {
	// Write your code here
	chars := []rune(s)
	changes := make([]bool, len(chars)) // To track which positions were changed
	left, right := 0, len(chars)-1

	// Step 1: Make the string a palindrome
	for left < right {
		if chars[left] != chars[right] {
			larger := max(chars[left], chars[right])
			chars[left], chars[right] = larger, larger
			changes[left], changes[right] = true, true
			k--
		}
		left++
		right--
	}

	// If k < 0, it was not possible to make the string a palindrome
	if k < 0 {
		return "-1"
	}

	// Step 2: Maximize the palindrome value
	left, right = 0, len(chars)-1
	for left <= right {
		if left == right { // Handle the middle character in odd-length strings
			if k > 0 {
				chars[left] = '9'
			}
		} else {
			if chars[left] < '9' {
				// If one of these positions was already changed, use 1 more change
				if changes[left] || changes[right] {
					if k >= 1 {
						chars[left], chars[right] = '9', '9'
						k--
					}
				} else {
					// If neither position was changed, it costs 2 changes
					if k >= 2 {
						chars[left], chars[right] = '9', '9'
						k -= 2
					}
				}
			}
		}
		left++
		right--
	}

	return string(chars)
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	firstMultipleInput := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	nTemp, err := strconv.ParseInt(firstMultipleInput[0], 10, 64)
	checkError(err)
	n := int32(nTemp)

	kTemp, err := strconv.ParseInt(firstMultipleInput[1], 10, 64)
	checkError(err)
	k := int32(kTemp)

	s := readLine(reader)

	result := highestValuePalindrome(s, n, k)

	fmt.Fprintf(writer, "%s\n", result)

	writer.Flush()
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
