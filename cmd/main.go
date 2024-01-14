package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

func main() {

	if len(os.Args) < 3 {
		fmt.Println("Usage: <program> <file-path> <metadata>")
		os.Exit(1) // Exit the program with a non-zero status
	}

	filepath := os.Args[1]
	metadata := os.Args[2]

	file, err := os.Open(filepath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error opening file: %v\n", err)
	}
	defer file.Close()

	data := make(map[string]map[string]string)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		data[line] = map[string]string{
			"metadata": metadata,
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Error reading file: %v\n", err)
		os.Exit(1)
	}

	jsonData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error marshalling data: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(string(jsonData))
}
