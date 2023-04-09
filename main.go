package main

import (
	"github-scanner/dataprocesser"
	"os"
)

func getArgs() []string {
	inputArgs := os.Args[1:]
	if len(inputArgs) == 0 {
		panic("No arguments")
	}

	return inputArgs
}

func main() {
	dataprocesser.ProcessAccounts(getArgs())
}
