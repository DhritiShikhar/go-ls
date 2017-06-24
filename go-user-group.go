// print the user id and the group id and other details

package main

import (
	"fmt"
	"os/user"
)

func main() {
	userName := "dhriti"
	ThisUser, _ := user.Lookup(userName)
	fmt.Println("User Name:", ThisUser.Name)
	fmt.Println("User ID:", ThisUser.Uid)
	fmt.Println("Group ID:", ThisUser.Gid)
	fmt.Println("Group Name:", ThisUser.Username)
	fmt.Println("Home Dir:", ThisUser.HomeDir)
}
