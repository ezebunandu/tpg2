package main

import (
	"fmt"
	"os"
    "bufio"
    "strings"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: go run main.go <filepath> <match string>")
		os.Exit(1)
	}

	filepath := os.Args[1]
	matchString := os.Args[2]

	file, err := os.Open(filepath)
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, matchString) {
			fmt.Println(line)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading file: %v\n", err)
	}
}
