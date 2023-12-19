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

		// From the front, find the first number
		for _, char := range code {
			char := string(char)

			if isNumber(char) {
				firstDigit = char
				break
			}
		}

		// From the back, find the second number

		for i := len(code) - 1; i >= 0; i-- {
			char := string(code[i])
			if isNumber(char) {
				secondDigit = char
				break
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

func isNumber(char string) bool {
	if char == "0" || char == "1" || char == "2" || char == "3" || char == "4" || char == "5" || char == "6" || char == "7" || char == "8" || char == "9" {
		return true
	}

	return false
}
