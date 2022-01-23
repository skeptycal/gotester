package gotester

import (
	"path/filepath"
)

// find all go files
// parse go files
// - look for functions
// - identify inputs and outputs
// - generate random inputs
// - allow for user defined inputs
// - test and verify outputs
// - write tests and benchmarks
// - write examples for output functions

func getGoFiles() []string {
	matches, err := filepath.Glob("**/*.go")
	die(err)

	return matches
}

type (
	loc int // offset of function within file

	funcList struct {
		filename string
		funcMap  map[string]loc
	}
)

func parseGoFiles(files []string) {}
