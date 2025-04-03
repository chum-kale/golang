//env variables - secrte config data which are retrived by app when runnning

package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	//setting a key value pair
	os.Setenv("FOO", "1")
	fmt.Println("FOO:", os.Getenv("FOO"))
	fmt.Println("BAR:", os.Getenv("BAR"))

	//list all key-value
	fmt.Println()
	for _, e := range os.Environ() {
		pair := strings.SplitN(e, "=", 2)
		fmt.Println(pair[0])
	}
}
