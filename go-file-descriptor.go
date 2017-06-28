// get file descriptor

package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {
	// input filename as cli flag
	fileName := flag.String("file", "", "file name")
	flag.Parse()

	// open file
	filePointer, err := os.Open(*fileName)
	if err != nil {
		log.Fatal(err)
	}

	// get file descriptor
	fileDes := filePointer.Fd()
	fmt.Println(fileDes)

	filePointer2, err := os.Open("gols.go")
	if err != nil {
		log.Fatal(err)
	}

	// get file descriptor
	fileDes2 := filePointer2.Fd()
	fmt.Println(fileDes2)
}
