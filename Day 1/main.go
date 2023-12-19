package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	total := 0

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
	for _, code := range codes {

		firstDigit := ""
		secondDigit := ""

		// Find each number
		for _, char := range code {
			char := string(char)

			if charIsNumber(char) {
				if firstDigit == "" {
					firstDigit = char
				}
				secondDigit = char
			}
		}

		// Combine the digits as strings
		number := firstDigit + secondDigit

		// Cast to int
		num, err := strconv.Atoi(number)
		if err != nil {
			log.Fatal(err)
		}

		// Add to total
		total += num

	}
	fmt.Printf("Total: %d\n", total)
}

func charIsNumber(char string) bool {
	if char == "0" || char == "1" || char == "2" || char == "3" || char == "4" || char == "5" || char == "6" || char == "7" || char == "8" || char == "9" {
		return true
	}

	return false
}

// func checkForSpelledNumber(code string) bool {
// 	if char == "o" {

// 	}

// 	// Check for 'two', 'three'
// 	if char == "t" {

// 	}

// 	// Check for 'four', 'five'
// 	if char == "f" {

// 	}

// 	// Check for 'six', 'seven'
// 	if char == "s" {

// 	}

// 	// Check for 'eight'
// 	if char == "e" {

// 	}

// 	// Check for 'nine'
// 	if char == "n" {

// 	}

// 	return false
// }
