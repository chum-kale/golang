package main

import (
	"bufio"
	"fmt"
	"os"
)

// check for error
func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	//dumping into a file without opening
	d1 := []byte("hello\ngo\n")
	err := os.WriteFile("/tmp/dat1", d1, 0644)
	check(err)

	//generally, open a file to write
	f, err := os.Create("/tmp/dat2")
	check(err)

	//close using defer
	defer f.Close()

	//writing byte slices
	d2 := []byte{115, 111, 109, 101, 10}
	n2, err := f.Write(d2)
	check(err)
	fmt.Printf("wrote %d bytes\n", n2)

	//write string
	n3, err := f.WriteString("writes\n")
	check(err)
	fmt.Printf("wrote %d bytes\n", n3)

	//sync to flush writes to stable storage
	f.Sync()

	//buffered writers
	w := bufio.NewWriter(f)
	n4, err := w.WriteString("buffered\n")
	check(err)
	fmt.Printf("wrote %d bytes\n", n4)

	//flush ensures all buffered ops have been done
	w.Flush()
}
