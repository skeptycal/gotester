package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {

	pattern := "*.*"

	if len(os.Args) > 1 {
		pattern = os.Args[1]
	}

	matches, err := filepath.Glob(pattern)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(matches)

	// search upper directory
	upperDirPattern := "../" + pattern

	// you can specify directly the directory you want to search as well with the pattern
	// for example, /usr/files/*input*

	matches, err = filepath.Glob(upperDirPattern)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(matches)
}
