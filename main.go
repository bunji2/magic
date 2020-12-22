package main

import (
	"fmt"
	"os"
)

const (
	usageFmt  = "Usage: %s file"
	magicFile = "magic.json"
)

func main() {
	os.Exit(run())
}

func run() (exitCode int) {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, usageFmt, os.Args[0])
		exitCode = 1
		return
	}
	filePath := os.Args[1]

	err := loadMagic(magicFile)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		exitCode = 2
		return
	}

	err = processFile(filePath)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		exitCode = 3
	}

	return
}
