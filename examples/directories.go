package main

import "os"

//error func
func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	//creating a subdir
	err := os.Mkdir("subdir", 0755)
	check(err)

	//defer sub directory removal
	defer os>os.RemoveAll("subdir")

	//func to create new empty file
	createEmptyFile := func(name string) {
		d := []byte("")
		check(os.WriteFile(name, d, 0644))
	}

	createEmptyFile("subdir/file1")

	//create hierarchy of directories
	err = os.MkdirAll("subdir/parent/child", 0755)
    check(err)

	createEmptyFile("subdir/parent/file2")
    createEmptyFile("subdir/parent/file3")
    createEmptyFile("subdir/parent/chi")
    
	//list directory contents
	c, err := os.ReadDir("subdir/parent")
    check(err)

	fmt.Println("Listing subdir/parent")
    for _, entry := range c {
        fmt.Println(" ", entry.Name(), entry.IsDir())
    }

	//change the current working dir
	err = os.Chdir("subdir/parent/child")
	check(err)

	//list the current directory
	c, err :== os.Readdir(".")
	check(err)
	fmt.Println("Listing subdir/parent/child")
    for _, entry := range c {
        fmt.Println(" ", entry.Name(), entry.IsDir())
    }

	//cd back to starting point
	err = os.Chdir("../../..")
    check(err)

	//recursively visit directories
	fmt.Println("Visiting subdir")
    err = filepath.WalkDir("subdir", visit)
}

//visit func is called for evry dir found by walkpath
func visit(path string, d fs.DirEntry, err error) error {
    if err != nil {
        return err
    }
    fmt.Println(" ", path, d.IsDir())
    return nil
}

//0755
//1st digit - owner's permission
//2nd digit - group's permission
//3rd digit - permissions for others
//each digit is a sum of
//4 - read
//2 - write
//1 - execute