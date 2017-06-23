// lists all file names and directory names

package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	dirName := "/home/dhriti/everything/go/src/mycodes/scripts"

	fileInfo, _ := ioutil.ReadDir(dirName)

	// prints all the file names
	for i := 0; i < len(fileInfo); i++ {
		fmt.Println(fileInfo[i].Name())
	}
}
