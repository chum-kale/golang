package main

import (
	"flag"
	"fmt"
)

func main() {
	//declare a flag
	//returns a string pointer
	wordPtr := flag.String("word", "foo", "a string")

	//num and fork flags
	numbPtr := flag.Int("numb", 42, "an int")
	forkPtr := flag.Bool("fork", false, "a bool")

	//use existing var
	var svar string
	flag.StringVar(&svar, "svar", "bar", "a string var")

	//call aftre declaring all flags
	flag.Parse()

	fmt.Println("word:", *wordPtr)
	fmt.Println("numb:", *numbPtr)
	fmt.Println("fork:", *forkPtr)
	fmt.Println("svar:", svar)
	fmt.Println("tail:", flag.Args())
}
