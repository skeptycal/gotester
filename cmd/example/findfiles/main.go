package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/skeptycal/gotester"
)

var (
	flagSubDir = *flag.Bool("subdir", false, "search in subdirectories")
)

func init() {
	flag.Parse()
}

func main() {

	pattern := "*.*"

	if len(os.Args) > 1 {
		pattern = os.Args[1]
	}

	a, err := filepath.Abs(pattern)
	gotester.Die(err)

	base := Base(a)
	fmt.Println("base: ", base)

	dir := Dir(a)
	fmt.Println("dir: ", dir)

	fmt.Println()

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
