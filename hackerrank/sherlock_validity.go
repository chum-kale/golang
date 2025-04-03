package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

/*
 * Complete the 'isValid' function below.
 *
 * The function is expected to return a STRING.
 * The function accepts STRING s as parameter.
 */

func isValid(s string) string {
	// Write your code here
	char_map := make(map[rune]int)
	for _, char := range s {
		char_map[char]++
	}

	// Determine the frequency of frequencies
	var freq1, freq2, count1, count2 int
	for _, freq := range char_map {
		if freq1 == 0 || freq == freq1 {
			freq1 = freq
			count1++
		} else if freq2 == 0 || freq == freq2 {
			freq2 = freq
			count2++
		} else {
			return "NO"
		}
	}

	// If there is only one frequency, it's valid
	if freq2 == 0 {
		return "YES"
	}

	// Check validity conditions
	if (count1 == 1 && (freq1-1 == freq2 || freq1-1 == 0)) ||
		(count2 == 1 && (freq2-1 == freq1 || freq2-1 == 0)) {
		return "YES"
	}

	return "NO"
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	s := readLine(reader)

	result := isValid(s)

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

// other way
// char_map := make(map[rune]int)

// for _, char := range s {
// 	char_map[char] += 1
// }

// freq_map := make(map[int]int)

// for _, freq := range char_map {
// 	freq_map[freq] += 1
// }
// 	if len(freq_map) == 1 {
// 	return "YES"
// }

// // If there are more than two different frequencies, it's invalid
// if len(freq_map) > 2 {
// 	return "NO"
// }

// // If there are exactly two different frequencies, check further
// freqList := []int{}
// counts := []int{}
// for freq, count := range freq_map {
// 	freqList = append(freqList, freq)
// 	counts = append(counts, count)
// }

// // The two frequencies and their counts
// freq1, freq2 := freqList[0], freqList[1]
// count1, count2 := counts[0], counts[1]

// // Check if it's valid by removing one character
// // Case 1: One frequency is 1, and it occurs only once (e.g., 1x1, 2x5)
// if (freq1 == 1 && count1 == 1) || (freq2 == 1 && count2 == 1) {
// 	return "YES"
// }

// // Case 2: Difference between the two frequencies is 1, and the higher frequency occurs only once
// if (freq1-freq2 == 1 && count1 == 1) || (freq2-freq1 == 1 && count2 == 1) {
// 	return "YES"
// }

// return "NO"
