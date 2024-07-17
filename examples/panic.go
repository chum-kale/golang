//panic - used for danger errors, that can't be handled gracefully

package main

import "os"

func main() {
	panic("a problem")

	//panic when a func returns an error value we can't handle
	//err when creating a new file
	_, err := os.Create("/tmp/file")
	if err != nil {
		panic(err)
	}
}
