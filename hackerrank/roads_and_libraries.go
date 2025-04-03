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
 * Complete the 'roadsAndLibraries' function below.
 *
 * The function is expected to return a LONG_INTEGER.
 * The function accepts following parameters:
 *  1. INTEGER n
 *  2. INTEGER c_lib
 *  3. INTEGER c_road
 *  4. 2D_INTEGER_ARRAY cities
 */

func roadsAndLibraries(n int32, c_lib int32, c_road int32, cities [][]int32) int64 {
	// If libraries are cheaper than roads, build library in each city
	if c_lib <= c_road {
		return int64(n) * int64(c_lib)
	}

	// Initialize disjoint set
	parent := make([]int32, n+1)
	size := make([]int32, n+1)
	for i := int32(1); i <= n; i++ {
		parent[i] = i
		size[i] = 1
	}

	// Find function for disjoint set with path compression
	var find func(int32) int32
	find = func(x int32) int32 {
		if parent[x] != x {
			parent[x] = find(parent[x])
		}
		return parent[x]
	}

	// Union function with size balancing
	union := func(x, y int32) {
		px, py := find(x), find(y)
		if px != py {
			if size[px] < size[py] {
				px, py = py, px
			}
			parent[py] = px
			size[px] += size[py]
		}
	}

	// Connect cities using roads
	for _, city := range cities {
		union(city[0], city[1])
	}

	// Count number of components and their sizes
	components := make(map[int32]int32)
	for i := int32(1); i <= n; i++ {
		root := find(i)
		components[root]++
	}

	// Calculate minimum cost
	// For each component:
	// - Need 1 library
	// - Need (size-1) roads to connect all cities
	var totalCost int64
	for _, componentSize := range components {
		roads := componentSize - 1
		totalCost += int64(c_lib) + int64(roads)*int64(c_road)
	}

	return totalCost
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	qTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	q := int32(qTemp)

	for qItr := 0; qItr < int(q); qItr++ {
		firstMultipleInput := strings.Split(strings.TrimSpace(readLine(reader)), " ")

		nTemp, err := strconv.ParseInt(firstMultipleInput[0], 10, 64)
		checkError(err)
		n := int32(nTemp)

		mTemp, err := strconv.ParseInt(firstMultipleInput[1], 10, 64)
		checkError(err)
		m := int32(mTemp)

		c_libTemp, err := strconv.ParseInt(firstMultipleInput[2], 10, 64)
		checkError(err)
		c_lib := int32(c_libTemp)

		c_roadTemp, err := strconv.ParseInt(firstMultipleInput[3], 10, 64)
		checkError(err)
		c_road := int32(c_roadTemp)

		var cities [][]int32
		for i := 0; i < int(m); i++ {
			citiesRowTemp := strings.Split(strings.TrimRight(readLine(reader), " \t\r\n"), " ")

			var citiesRow []int32
			for _, citiesRowItem := range citiesRowTemp {
				citiesItemTemp, err := strconv.ParseInt(citiesRowItem, 10, 64)
				checkError(err)
				citiesItem := int32(citiesItemTemp)
				citiesRow = append(citiesRow, citiesItem)
			}

			if len(citiesRow) != 2 {
				panic("Bad input")
			}

			cities = append(cities, citiesRow)
		}

		result := roadsAndLibraries(n, c_lib, c_road, cities)

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
