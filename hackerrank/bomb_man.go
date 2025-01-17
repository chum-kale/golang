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
 * Complete the 'bomberMan' function below.
 *
 * The function is expected to return a STRING_ARRAY.
 * The function accepts following parameters:
 *  1. INTEGER n
 *  2. STRING_ARRAY grid
 */

func bomberMan(n int32, grid []string) []string {
	// Write your code here
	rows := len(grid)
	cols := len(grid[0])

	detonate := func(currentGrid []string) []string {
		newGrid := make([][]rune, rows)
		for i := 0; i < rows; i++ {
			newGrid[i] = make([]rune, cols)
			for j := 0; j < cols; j++ {
				newGrid[i][j] = 'o'
			}
		}
		for r := 0; r < rows; r++ {
			for c := 0; c < cols; c++ {
				if currentGrid[r][c] == 'o' {
					newGrid[r][c] = '.'
					if r > 0 {
						newGrid[r-1][c] = '.'
					}
					if r < rows-1 {
						newGrid[r+1][c] = '.'
					}
					if c > 0 {
						newGrid[r][c-1] = '.'
					}
					if c < cols-1 {
						newGrid[r][c+1] = '.'
					}
				}
			}
		}
		res := make([]string, rows)
		for i := 0; i < rows; i++ {
			res[i] = string(newGrid[i])
		}
		return res
	}

	if n == 1 {
		return grid
	}

	if n%2 == 0 {
		full := strings.Repeat("o", cols)
		result := make([]string, rows)
		for i := range result {
			result[i] = full
		}
		return result
	}

	state3 := detonate(grid)
	state5 := detonate(state3)

	if (n-3)%4 == 0 {
		return state3
	}
	return state5
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	firstMultipleInput := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	rTemp, err := strconv.ParseInt(firstMultipleInput[0], 10, 64)
	checkError(err)
	r := int32(rTemp)

	cTemp, err := strconv.ParseInt(firstMultipleInput[1], 10, 64)
	checkError(err)
	c := int32(cTemp)

	nTemp, err := strconv.ParseInt(firstMultipleInput[2], 10, 64)
	checkError(err)
	n := int32(nTemp)

	var grid []string

	for i := 0; i < int(r); i++ {
		gridItem := readLine(reader)
		grid = append(grid, gridItem)
	}

	result := bomberMan(n, grid)

	for i, resultItem := range result {
		fmt.Fprintf(writer, "%s", resultItem)

		if i != len(result)-1 {
			fmt.Fprintf(writer, "\n")
		}
	}

	fmt.Fprintf(writer, "\n")

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
