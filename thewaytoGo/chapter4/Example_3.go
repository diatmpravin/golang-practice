// Example_3.go
package main

import (
	"fmt"
	"os/user"
)

func main() {
	currentUser, err := user.Current()

	if err != nil {
		return
	}

	configDir := currentUser.HomeDir
	file := configDir + "/test.rb"

	if err := file.Chmod(0664); err != nil {
		fmt.Println(err)
		return err
	}
}
