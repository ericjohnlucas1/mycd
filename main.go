package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Exactly two arguments are required. The first argument should be the current directory, and the second argument should be the destination.")
		os.Exit(0)
	}
	curdir, newdir := os.Args[1], os.Args[2]

	destpath, err := NewPath(newdir)
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
	curpath, err := NewPath(curdir)
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}

	if destpath.isAbsolute() {
		destpath.shorten()
		fmt.Println(destpath.getStringRepresentation())

	} else {
		resultpath := curpath.concatenate(destpath)
		resultpath.shorten()
		fmt.Println(resultpath.getStringRepresentation())
	}
}
