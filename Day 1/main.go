package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	// Import text file
    input, err := os.Open("input.txt")
    if err != nil {
        log.Fatal(err)
    }
    defer input.Close()

    scanner := bufio.NewScanner(input)
    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }

	var codes []string

	// Put each line into an array
	for scanner.Scan() {
		codes = append(codes, scanner.Text())
    }

	// Loop through array
	for i, code := range codes {
		fmt.Println(i, code)
	}





	// From the front, find the first number

	// From the back, find the second number

	// Combine the digits as strings

	// Cast to int

	// Add to total

	// Return total

}
