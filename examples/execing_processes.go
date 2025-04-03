//replace current go process with another one
//can be non-go

package main

import (
	"os"
	"os/exec"
	"syscall"
)

func main() {
	//abs path to binaryy we want to execute
	binary, lookErr := exec.LookPath("ls")
	if lookErr != nil {
		panic(lookErr)
	}

	//execr requires args in slice form
	args := []string{"ls", "-a", "-l", "-h"}

	//exec needs env variables
	//provide current env here
	env := os.Environ()

	//actual call
	//if successful, it will get replaced by specified binary
	execErr := syscall.Exec(binary, args, env)
	if execErr != nil {
		panic(execErr)
	}
}
