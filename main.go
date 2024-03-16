package main

import (
	"fmt"
	"io/fs"
	"log"
	"os"
)

type fileInfo struct {
	path string
	fs.DirEntry
}

func (f *fileInfo) FullPath() string {
	return f.path
}

func main() {
	args := os.Args

	if len(args) < 3 {
		fmt.Println("Usage: go run main.go <pattern> <path>")
		os.Exit(1)
	}

	pattern := args[1]
	path := args[2]

	log.Printf("Pattern: %s, Path: %s", pattern, path)

	files, err := readDirRecursively(path)
	if err != nil {
		log.Fatal(err)
	}

	matchedDirs := matchpattern(files, pattern)

	for _, f := range matchedDirs {
		fmt.Println(f.FullPath())
	}
}

func matchpattern(files []fileInfo, pattern string) []fileInfo {
	fmt.Println("Pattern: ", pattern)
	return files
}

func readDirRecursively(path string) ([]fileInfo, error) {
	dirs, err := os.ReadDir(path)
	if err != nil {
		return nil, err
	}

	var entries []fileInfo
	for _, d := range dirs {
		if d.IsDir() {
			subDir, err := readDirRecursively(path + "/" + d.Name())
			if err != nil {
				return nil, err
			}
			entries = append(entries, subDir...)
		} else {
			f := fileInfo{path: path + "/" + d.Name(), DirEntry: d}
			entries = append(entries, f)
		}
	}

	return entries, nil
}
