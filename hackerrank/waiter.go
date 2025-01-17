// chatgpt soln
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
 * Complete the 'waiter' function below.
 *
 * The function is expected to return an INTEGER_ARRAY.
 * The function accepts following parameters:
 *  1. INTEGER_ARRAY number
 *  2. INTEGER q
 */

func getPrimes(n int32) []int32 {
	primes := []int32{}
	num := int32(2)
	for int32(len(primes)) < n {
		isPrime := true
		for i := int32(2); i*i <= num; i++ {
			if num%i == 0 {
				isPrime = false
				break
			}
		}
		if isPrime {
			primes = append(primes, num)
		}
		num++
	}
	return primes
}

func waiter(number []int32, q int32) []int32 {
	// Write your code here
	primes := getPrimes(q)
	result := []int32{}
	currentStack := number

	for i := int32(0); i < q; i++ {
		prime := primes[i]
		A := []int32{}
		B := []int32{}

		// Separate plates into A and B
		for len(currentStack) > 0 {
			plate := currentStack[len(currentStack)-1] // Pop from the top
			currentStack = currentStack[:len(currentStack)-1]
			if plate%prime == 0 {
				A = append(A, plate)
			} else {
				B = append(B, plate)
			}
		}

		// Append A to the result in the reverse order
		for len(A) > 0 {
			result = append(result, A[len(A)-1])
			A = A[:len(A)-1]
		}

		// Set B as the new stack for the next iteration
		currentStack = B
	}

	// Append remaining plates in B to the result
	for len(currentStack) > 0 {
		result = append(result, currentStack[len(currentStack)-1])
		currentStack = currentStack[:len(currentStack)-1]
	}

	return result
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

	qTemp, err := strconv.ParseInt(firstMultipleInput[1], 10, 64)
	checkError(err)
	q := int32(qTemp)

	numberTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	var number []int32

	for i := 0; i < int(n); i++ {
		numberItemTemp, err := strconv.ParseInt(numberTemp[i], 10, 64)
		checkError(err)
		numberItem := int32(numberItemTemp)
		number = append(number, numberItem)
	}

	result := waiter(number, q)

	for i, resultItem := range result {
		fmt.Fprintf(writer, "%d", resultItem)

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
