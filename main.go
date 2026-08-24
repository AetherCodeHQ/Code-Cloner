package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	src := "."
	if len(os.Args) > 1 {
		src = os.Args[1]
	}
	fmt.Println("Code Cloner")
	fmt.Printf("Scanning %s for source files...\n", src)
	exts := map[string]int{".go": 0, ".py": 0, ".js": 0, ".ts": 0, ".rs": 0, ".java": 0}
	var totalLines int64
	var files []string
	filepath.Walk(src, func(path string, info os.FileInfo, err error) error {
		if err != nil || info.IsDir() {
			return nil
		}
		ext := filepath.Ext(path)
		if _, ok := exts[ext]; ok {
			exts[ext]++
			totalLines += info.Size() / 50
			files = append(files, path)
		}
		return nil
	})
	fmt.Println("\nFiles found by language:")
	for ext, count := range exts {
		if count > 0 {
			fmt.Printf("  %s: %d files\n", ext, count)
		}
	}
	fmt.Printf("\nTotal: %d source files, ~%d lines\n", len(files), totalLines)
}
