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

	for scanner.Scan() {
        fmt.Println(scanner.Text())
    }


	// Put each line into an array

	// Loop through array

	// From the front, find the first number

	// From the back, find the second number

	// Combine the digits as strings

	// Cast to int

	// Add to total

	// Return total

}
