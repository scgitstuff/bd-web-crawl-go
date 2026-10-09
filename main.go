package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	// fmt.Println("Hello, World!")

	if len(os.Args) < 2 {
		fmt.Println("no website provided")
		os.Exit(1)
	}

	if len(os.Args) > 2 {
		fmt.Println("too many arguments provided")
		os.Exit(1)
	}

	BASE_URL := strings.TrimSpace(os.Args[1])
	fmt.Printf("\nstarting crawl of: %s\n", BASE_URL)
}
