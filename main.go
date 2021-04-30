package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/AsyncEngineering/nomad"
)

func main() {
	usr, err := user.Current()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Hello %s, welcome to Nomad ⛺\n", usr.Username)
	repl.Start(os.Stdin, os.Stdout)
}
