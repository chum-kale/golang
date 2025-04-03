package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

// check for errors
func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	//slurping file's content
	dat, err := os.ReadFile("/tmp/dat")
	check(err)
	fmt.Print(string(dat))

	//obtain os.file value
	f, err := os.Open("/tmp/dat")
	check(err)

	//read bytes from start (5)
	b1 := make([]byte, 5)
	n1, err := f.Read(b1)
	check(err)
	fmt.Printf("%d bytes: %s\n", n1, string(b1[:n1]))

	//reading file from a known location
	o2, err := f.Seek(6, io.SeekStart)
	check(err)
	b2 := make([]byte, 2)
	n2, err := f.Read(b2)
	check(err)
	fmt.Printf("%d bytes @ %d: ", n2, o2)
	fmt.Printf("%v\n", string(b2[:n2]))

	//eeking methods relative to current cusrosr position
	_, err = f.Seek(4, io.SeekCurrent)
	check(err)

	//relative to end of file
	_, err = f.Seek(-10, io.SeekEnd)
	check(err)

	//io package
	o3, err := f.Seek(6, io.SeekStart)
	check(err)
	b3 := make([]byte, 2)
	n3, err := io.ReadAtLeast(f, b3, 2)
	check(err)
	fmt.Printf("%d bytes @ %d: %s\n", n3, o3, string(b3))

	//rewind
	_, err = f.Seek(0, io.SeekStart)
	check(err)

	//buffered reader
	r4 := bufio.NewReader(f)
	b4, err := r4.Peek(5)
	check(err)
	fmt.Printf("5 bytes: %s\n", string(b4))

	//close file should be done with defer
	f.Close()
}
