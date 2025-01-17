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
 * Complete the 'sumXor' function below.
 *
 * The function is expected to return a LONG_INTEGER.
 * The function accepts LONG_INTEGER n as parameter.
 */

func sumXor1(n int64) int64 {
	// Write your code here
	if n == 0 {
		return 1
	}

	var count int64 = 0
	for i := 0; i < int(n); i++ {
		if int(n)+i == int(n)^i {
			count += 1
		}
	}

	return count
}

func sumXor(n int64) int64 {
	// If n is zero, there is one value x that satisfies the equation: x = 0
	if n == 0 {
		return 1
	}

	// Count the number of zero bits in n (i.e., the number of 0s in binary representation)
	var count int64 = 0
	for n > 0 {
		if n&1 == 0 {
			count++
		}
		n >>= 1
	}

	// For each zero bit, there are two possibilities for x (either 0 or 1 at that position)
	return 1 << count // 2^count
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	n, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)

	result := sumXor(n)

	fmt.Fprintf(writer, "%d\n", result)

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
