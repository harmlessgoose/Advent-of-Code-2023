package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	// Import input file
	input, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer input.Close()

	scanner := bufio.NewScanner(input)
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	var games []string

	// Put each line into an array
	for scanner.Scan() {
		games = append(games, scanner.Text())
	}

	// Loop through array
	for _, game := range games {

		fmt.Println(game)

		// For each line, scan for each word

		// If the word is a colour, extract the number

		// Add the number to a running total

		// Based on the total, decide if the game is valid

		// If so, extract the game ID and add to the total

	}
	// Return the total
}
