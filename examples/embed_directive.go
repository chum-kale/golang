// //go:embed is a compiler directive that allows programs to include arbitrary files and folders in the Go binary at build time
package main

import "embed"

//accepts path relative to directory containing go source file
//embeds contents of file into string immediately after it

//go:embed folder/single_file.txt
var fileString string

//embed file contents into byte

//go:embed folder/single_file.txt
var fileByte []byte

//virtual file system

//go:embed folder/single_file.txt
//go:embed folder/*.hash
var folder embed.FS

func main() {
	//Print out the contents of single_file.txt.

	print(fileString)
	print(string(fileByte))
	//Retrieve some files from the embedded folder.

	content1, _ := folder.ReadFile("folder/file1.hash")
	print(string(content1))
	content2, _ := folder.ReadFile("folder/file2.hash")
	print(string(content2))
}
