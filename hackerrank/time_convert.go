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
 * Complete the 'timeConversion' function below.
 *
 * The function is expected to return a STRING.
 * The function accepts STRING s as parameter.
 */

func timeConversion(s string) string {
	// Write your code here
	//s = strings.ReplaceAll(s, "PM", "")
	//s = strings.ReplaceAll(s, "AM", "")
	s = strings.TrimSpace(s)
	parts := strings.Split(s, ":")

	hour, err := strconv.Atoi(parts[0])
	if err != nil {
		fmt.Println("invalid value")
	}

	if strings.Contains(s, "AM") {
		if hour == 12 {
			parts[0] = "00"
			parts[2] = strings.ReplaceAll(parts[2], "AM", "")
			return strings.Join(parts, ":")
		} else {
			s = strings.ReplaceAll(s, "AM", "")
			return s
		}
	} else {
		if hour == 12 {
			s = strings.ReplaceAll(s, "PM", "")
			return s
		} else {
			hour += 12
			parts[0] = fmt.Sprintf("%02d", hour)
			parts[2] = strings.ReplaceAll(parts[2], "PM", "")
			return strings.Join(parts, ":")
		}
	}
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	s := readLine(reader)

	result := timeConversion(s)

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
