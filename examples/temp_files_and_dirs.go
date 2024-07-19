package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	//creating temporary file
	//"" is first srg meaning, file will be created at default loc
	//returns open file
	f, err := os.CreateTemp("", "sample")
	check(err)

	//display name of temp file
	fmt.Println("Temp file name:", f.Name())

	//cleanup
	defer os.Remove(f.Name())

	//write data into file
	_, err = f.Write([]byte{1, 2, 3, 4})
	check(err)

	//creating temporary dir
	//returns dir name rather than open file
	dname, err := os.MkdirTemp("", "sampledir")
	check(err)
	fmt.Println("Temp dir name:", dname)

	defer os.RemoveAll(dname)

	//syntehsize temp file names by prefixing them with temp dir
	fname := filepath.Join(dname, "file1")
	err = os.WriteFile(fname, []byte{1, 2}, 0666)
	check(err)
}
