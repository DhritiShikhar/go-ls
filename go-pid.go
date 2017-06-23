// Get pid of the calling process and parent process

package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Process ID: ", os.Getpid())
	fmt.Println("Parent process ID: ", os.Getppid())
}
