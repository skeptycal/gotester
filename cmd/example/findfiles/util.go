package main

import (
	"os"
	"path/filepath"
)

func Base(path string) (s string) {
	_, s = filepath.Split(path)
	// filepath.Base(s)
	return
}

func Dir(path string) (s string) {
	s, _ = filepath.Split(path)
	return
}

func Exists(file string) bool {
	_, err := os.Stat(file)
	return err == nil
}

func Mode(file string) string {
	info, err := os.Stat(file)
	if err != nil {
		// return ""
		return "- error! -"
	}
	return info.Mode().String()
}
