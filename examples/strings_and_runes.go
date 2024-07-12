package main

import (
	"fmt"
	"unicode/utf8"
)

//string - readonly slice of bytes
//containers of text encoded in utf-8
//rune - integer that represnts unicode code point (similar to character)

func main() {
	const s = "สวัสดี"

	//string = []byte
	fmt.Println("Len:", len(s))

	//indexing a string produces raw byte values at each index
	for i := 0; i < len(s); i++ {
		fmt.Printf("%x ", s[i])
	}
	fmt.Println()

	fmt.Println("Rune count:", utf8.RuneCountInString(s))

	for idx, runeValue := range s {
		fmt.Printf("%#U starts at %d\n", runeValue, idx)
	}

	//same iteration uising functions
	fmt.Println("\nUsing DecodeRuneInString")
	for i, w := 0, 0; i < len(s); i += w {
		runeValue, width := utf8.DecodeRuneInString(s[i:])
		fmt.Printf("%#U starts at %d\n", runeValue, i)
		w = width
		examineRune(runeValue)
	}
}

func examineRune(r rune) {
	//Values enclosed in single quotes are rune literals. We can compare a rune value to a rune literal directly.

	if r == 't' {
		fmt.Println("found tee")
	} else if r == 'ส' {
		fmt.Println("found so sua")
	}
}
