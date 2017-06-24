// Read commands from standard input and execute them
// Implementation of shell-like program

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
)

func main() {
	// Read the command input from user
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter command>>> ")
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}

	command := exec.Command(input)
	out, err := command.Output()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(out)

}
