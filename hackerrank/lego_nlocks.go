// chatgpt version
// doesnt run

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
 * Complete the 'legoBlocks' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts following parameters:
 *  1. INTEGER n
 *  2. INTEGER m
 */

const MOD int32 = 1e9 + 7

// Helper function to calculate modular exponentiation
func modPow(base, exp int32) int32 {
	result := int32(1)
	for exp > 0 {
		if exp%2 == 1 {
			result = (result * base) % MOD
		}
		base = (base * base) % MOD
		exp /= 2
	}
	return result
}

func legoBlocks(n int32, m int32) int32 {
	// Step 1: Calculate the number of ways to fill a row
	rowWays := make([]int32, m+1)
	rowWays[0] = 1

	for i := int32(1); i <= m; i++ {
		rowWays[i] = rowWays[i-1]
		if i-2 >= 0 {
			rowWays[i] += rowWays[i-2]
		}
		if i-3 >= 0 {
			rowWays[i] += rowWays[i-3]
		}
		if i-4 >= 0 {
			rowWays[i] += rowWays[i-4]
		}
		rowWays[i] %= MOD
	}

	// Step 2: Calculate total ways to build a wall of height n
	totalWays := make([]int32, m+1)
	for i := int32(1); i <= m; i++ {
		totalWays[i] = modPow(rowWays[i], n)
	}

	// Step 3: Subtract invalid configurations to get valid ways
	validWays := make([]int32, m+1)
	for i := int32(1); i <= m; i++ {
		validWays[i] = totalWays[i]
		for j := int32(1); j < i; j++ {
			validWays[i] -= (validWays[j] * totalWays[i-j]) % MOD
			if validWays[i] < 0 {
				validWays[i] += MOD
			}
		}
		validWays[i] %= MOD
	}

	return validWays[m]
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	tTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	t := int32(tTemp)

	for tItr := 0; tItr < int(t); tItr++ {
		firstMultipleInput := strings.Split(strings.TrimSpace(readLine(reader)), " ")

		nTemp, err := strconv.ParseInt(firstMultipleInput[0], 10, 64)
		checkError(err)
		n := int32(nTemp)

		mTemp, err := strconv.ParseInt(firstMultipleInput[1], 10, 64)
		checkError(err)
		m := int32(mTemp)

		result := legoBlocks(n, m)

		fmt.Fprintf(writer, "%d\n", result)
	}

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
