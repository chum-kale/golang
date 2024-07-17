//defer - function call is performed later (ensure and finally)
//eg - cleanup

package main

import (
	"fmt"
	"os"
)

func main() {
	//create a file, write in it and close when done
	f := createFile("/tmp/defer.txt")
	//this defer will be executed at end once writeFile has finished
	defer closeFile(f)
	writeFile()
}

func createFile(p string) *os.File {
	fmt.Println("creating")
	f, err := os.Create(p)
	if err != nil {
		panic(err)
	}
	return f
}

func writeFile(f *os.File) {
	fmt.Println("writing")
	fmt.Fprintln(f, "data")
}

func closeFile(f *os.File) {
	fmt.Println("closing")
	err := f.close()

	//checking for errors when closing a fiel
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}
