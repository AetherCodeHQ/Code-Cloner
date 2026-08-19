package main

import (
	"fmt"
	"os"
)

// code_cloner - Clone code patterns
func code_cloner(path string) {
	fmt.Println("========================================")
	fmt.Println("  Code-Cloner")
	fmt.Println("  Clone code patterns")
	fmt.Println("========================================")
	fmt.Println()
	fmt.Println("Target:", path)
	fmt.Println("Processing...")
	fmt.Println("Done!")
}

func main() {
	path := "."
	if len(os.Args) > 1 {
		path = os.Args[1]
	}
	code_cloner(path)
}
