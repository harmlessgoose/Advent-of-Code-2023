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

		// Find each character
		for index, char := range code {
			char := string(char)
			number := ""

			// Check if char is a spelled number
			number = checkForSpelledNumber(code, index)

			// Check if char is a number
			if charIsNumber(char) {
				number = char
			}

			if number != "" {
				if firstDigit == "" {
					firstDigit = number
				}
				secondDigit = number
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

func checkForSpelledNumber(code string, index int) string {
	number := ""

	// Three letter numbers
	if isOutOfBounds(code, index, 3) {
		return ""
	}
	// Check for 'one'
	if code[index] == 'o' && code[index+1] == 'n' && code[index+2] == 'e' {
		number = "1"
	}
	// Check for 'two'
	if code[index] == 't' && code[index+1] == 'w' && code[index+2] == 'o' {
		number = "2"
	}
	// Check for 'six'
	if code[index] == 's' && code[index+1] == 'i' && code[index+2] == 'x' {
		number = "6"
	}

	if number != "" {
		return number
	}

	// Four letter numbers
	if isOutOfBounds(code, index, 4) {
		return ""
	}
	// Check for 'four'
	if code[index] == 'f' && code[index+1] == 'o' && code[index+2] == 'u' && code[index+3] == 'r' {
		number = "4"
	}

	// Check for 'five'
	if code[index] == 'f' && code[index+1] == 'i' && code[index+2] == 'v' && code[index+3] == 'e' {
		number = "5"
	}
	// Check for 'nine'
	if code[index] == 'n' && code[index+1] == 'i' && code[index+2] == 'n' && code[index+3] == 'e' {
		number = "9"
	}

	if number != "" {
		return number
	}

	// Five letter numbers
	if isOutOfBounds(code, index, 5) {
		return ""
	}

	// Check for 'three'
	if code[index] == 't' && code[index+1] == 'h' && code[index+2] == 'r' && code[index+3] == 'e' && code[index+4] == 'e' {
		number = "3"
	}

	// Check for 'seven'
	if code[index] == 's' && code[index+1] == 'e' && code[index+2] == 'v' && code[index+3] == 'e' && code[index+4] == 'n' {
		number = "7"
	}

	// Check for 'eight'
	if code[index] == 'e' && code[index+1] == 'i' && code[index+2] == 'g' && code[index+3] == 'h' && code[index+4] == 't' {
		number = "8"
	}

	return number
}

func isOutOfBounds(code string, index, lengthOfWord int) bool {
	return index+lengthOfWord > len(code)
}
