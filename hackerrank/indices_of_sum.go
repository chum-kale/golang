package main

import (
	"fmt"
)

func twoSum(nums []int, target int) []int {
	iMAp := make(map[int]int)

	for i, num := range nums {
		comp := target - num
		if idx, found := iMAp[comp]; found {
			return []int{idx, i}
		}

		iMAp[num] = i
	}

	return nil
}

func main() {
	arr := []int{2, 7, 11, 15}
	target := 9
	result := twoSum(arr, target)
	fmt.Println(result) // Output: [0, 1] (since 2 + 7 = 9)
}
