package main

import (
	"fmt"
	"one/repl"
	"os"
	"os/user"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	fmt.Printf(
		"Hello %s! Welcome to One!\n",
		user.Username)
	repl.Start(os.Stdin, os.Stdout)
}
