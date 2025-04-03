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
 * Complete the 'minimumMoves' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts following parameters:
 *  1. STRING_ARRAY grid
 *  2. INTEGER startX
 *  3. INTEGER startY
 *  4. INTEGER goalX
 *  5. INTEGER goalY
 */

type Position struct {
	x, y, dist int
}

func minimumMoves(grid []string, startX int32, startY int32, goalX int32, goalY int32) int32 {
	n := len(grid)

	// Create visited array to track positions and their distances
	visited := make([][]int, n)
	for i := range visited {
		visited[i] = make([]int, n)
		for j := range visited[i] {
			visited[i][j] = -1 // -1 means unvisited
		}
	}

	// Queue for BFS: [row, col, distance]
	type Position struct {
		x, y, dist int
	}
	queue := []Position{{int(startX), int(startY), 0}}
	visited[startX][startY] = 0

	// Directions: up, right, down, left
	directions := [][2]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}

	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]

		// If we reached the goal
		if curr.x == int(goalX) && curr.y == int(goalY) {
			return int32(curr.dist)
		}

		// Try each direction
		for _, dir := range directions {
			newX, newY := curr.x, curr.y

			// Keep moving in current direction until hitting wall/obstacle
			for {
				newX += dir[0]
				newY += dir[1]

				// Check if out of bounds
				if newX < 0 || newX >= n || newY < 0 || newY >= n {
					newX -= dir[0]
					newY -= dir[1]
					break
				}

				// Check if hit obstacle
				if grid[newX][newY] == 'X' {
					newX -= dir[0]
					newY -= dir[1]
					break
				}
			}

			// If new position is unvisited and different from current
			if visited[newX][newY] == -1 && (newX != curr.x || newY != curr.y) {
				visited[newX][newY] = curr.dist + 1
				queue = append(queue, Position{newX, newY, curr.dist + 1})
			}
		}
	}

	return -1 // Goal not reachable
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

	var grid []string

	for i := 0; i < int(n); i++ {
		gridItem := readLine(reader)
		grid = append(grid, gridItem)
	}

	firstMultipleInput := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	startXTemp, err := strconv.ParseInt(firstMultipleInput[0], 10, 64)
	checkError(err)
	startX := int32(startXTemp)

	startYTemp, err := strconv.ParseInt(firstMultipleInput[1], 10, 64)
	checkError(err)
	startY := int32(startYTemp)

	goalXTemp, err := strconv.ParseInt(firstMultipleInput[2], 10, 64)
	checkError(err)
	goalX := int32(goalXTemp)

	goalYTemp, err := strconv.ParseInt(firstMultipleInput[3], 10, 64)
	checkError(err)
	goalY := int32(goalYTemp)

	result := minimumMoves(grid, startX, startY, goalX, goalY)

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
