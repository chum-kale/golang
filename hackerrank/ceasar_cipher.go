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
 * Complete the 'caesarCipher' function below.
 *
 * The function is expected to return a STRING.
 * The function accepts following parameters:
 *  1. STRING s
 *  2. INTEGER k
 */

func caesarCipher(s string, k int32) string {
	// Write your code here
	var res strings.Builder

	// anything above 26 will be rounded off to 1-26
	k = k % 26

	for _, char := range s {
		if char >= 'a' && char <= 'z' {
			// char - a gives us absolute value of the alphabet
			// +k adds integer to that absolute value
			// %26 to round off integers to 1-26 then add a to get unicode.
			newch := (char-'a'+k)%26 + 'a'
			res.WriteRune(newch)
		} else if char >= 'A' && char <= 'Z' {
			newch := (char-'A'+k)%26 + 'A'
			res.WriteRune(newch)
		} else {
			res.WriteRune(char)
		}
	}

	return res.String()
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	n := int32(nTemp)

	s := readLine(reader)

	kTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	k := int32(kTemp)

	result := caesarCipher(s, k)

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
