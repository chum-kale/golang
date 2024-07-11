package main

import "fmt"

func main() {

	nums := []int{2, 3, 4}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum:", sum)

	for i, num := range nums {
		if num == 3 {
			fmt.Println("index:", i)
		}
	}

	//range can iterate over key value pairs
	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}

	//iterate over the keys only
	for k := range kvs {
		fmt.Println("key:", k)
	}

	//range in string iterates over unicode code pints
	//first value - starting byte index
	//second value - rune itself
	for i, c := range "go" {
		fmt.Println(i, c)
	}
}
