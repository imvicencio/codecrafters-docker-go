package main

import (
	"os"
	"os/exec"
)

// Usage: your_docker.sh run <image> <command> <arg1> <arg2> ...
func main() {
	// You can use print statements as follows for debugging, they'll be visible when running tests.

	command := os.Args[3]
	args := os.Args[4:len(os.Args)]
	
	cmd := exec.Command(command, args...)
	// output, err := cmd.Output()
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		os.Exit(1)
	}
	
	// fmt.Println(string(output))
}
