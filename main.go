package main

import (
	"os"

	"go-reloaded/functions"
)

func main() {
	args := os.Args[1:]

	if len(args) == 1 {
		functions.ProcessFile(args[0], args[0])
	} else if len(args) == 2 {
		functions.ProcessFile(args[0], args[1])
	}
}