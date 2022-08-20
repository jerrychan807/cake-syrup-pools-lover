package main

import (
	"cake-syrup-pools-lover/lib"
	"os"
)

func main() {
	defer os.Exit(0)
	cmd := lib.CommandLine{}
	cmd.Run()
}
