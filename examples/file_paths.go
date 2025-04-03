package main

import (
	"fmt"
	"path/filepath"
	"strings"
)

func main() {
	//constructing paths from a no opf arguments
	p := filepath.Join("dir1", "dir2", "filename")
	fmt.Println("p:", p)

	//don't concatenate , instead use join
	fmt.Println(filepath.Join("dir1//", "filename"))
	fmt.Println(filepath.Join("dir1/../dir1", "filename"))

	//dir - split path to directory
	//bse - split to file
	//split returns both in same call
	fmt.Println("Dir(p):", filepath.Dir(p))
	fmt.Println("Base(p):", filepath.Base(p))

	//check if path is absolute
	fmt.Println(filepath.IsAbs("dir/file"))
	fmt.Println(filepath.IsAbs("/dir/file"))

	filename := "config.json"

	//splitting out the extension
	ext := filepath.Ext(filename)
	fmt.Println(ext)

	//filenamew without extension
	fmt.Println(strings.TrimSuffix(filename, ext))

	//finding relative path between base and target, returns err if not found
	rel, err := filepath.Rel("a/b", "a/b/t/file")
	if err != nil {
		panic(err)
	}
	fmt.Println(rel)
	rel, err = filepath.Rel("a/b", "a/c/t/file")
	if err != nil {
		panic(err)
	}
	fmt.Println(rel)
}
