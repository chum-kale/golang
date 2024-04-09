package main

import "fmt"

func main() {
	//var colors map[string]string

	//colors := make(map[string]string)

	colors := map[string]string{
		"red":   "#ff000",
		"green": "#4f6785",
		"white": "#ffff",
	}

	//colors["white"] = "#fffff"

	//delete(colors, "white")

	fmt.Println((colors))
	printMap(colors)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("hex code for", color, "is", hex)
	}
}
