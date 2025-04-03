package main

import "fmt"

func reverseNumber(n int) int {
	reversed := 0
	for n != 0 {
		reversed = reversed*10 + n%10
		n /= 10
	}

	return reversed
}

func reverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i] // Swap characters
	}

	return string(runes)
}

func main() {
	fmt.Println(reverseNumber(1234))  // Output: 4321
	fmt.Println(reverseNumber(-9876)) // Output: -6789
	fmt.Println(reverseNumber(1000))  // Output: 1
}
