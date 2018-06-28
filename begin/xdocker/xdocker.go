package main

import (
	"os"
	"fmt"
	"os/exec"
)


func main() {
	switch os.Args[1] {
		case "run":
			run()
		default:
			panic("what??")
	}


}

func run() {
	fmt.Printf("Running process %v", os.Args[2:])
	cmd := exec.Command(os.Args[2], os.Args[3:]... )
	cmd.Stdin = os.Stdout
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()
}
