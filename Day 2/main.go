package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	var games []string
	idTotal := 0

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

	// Put each line from the input into an array
	for scanner.Scan() {
		games = append(games, scanner.Text())
	}

	// Loop through games
	for i, game := range games {

		// Split into subgames
		subGames := strings.Split(game, ";")

		highestRed := 0
		highestGreen := 0
		highestBlue := 0
		isPossible := true
		gameId := i + 1

		fmt.Printf("Game ID: %d\n", gameId)

		// For each subgame, extract each colour
		for _, subgame := range subGames {

			red := getColourValue(subgame, "red")
			if red > highestRed {
				highestRed = red
			}
			green := getColourValue(subgame, "green")
			if green > highestGreen {
				highestGreen = green
			}
			blue := getColourValue(subgame, "blue")
			if blue > highestBlue {
				highestBlue = blue
			}
		}

		// Based on the total, decide if the game is valid
		if highestRed > 12 || highestGreen > 13 || highestBlue > 14 {
			isPossible = false
		}

		if isPossible {
			idTotal += gameId
		}

		fmt.Printf("Highest Red is %d\n", highestRed)
		fmt.Printf("Highest Green is %d\n", highestGreen)
		fmt.Printf("Highest Blue is %d\n", highestBlue)
		fmt.Printf("Game is possible: %t\n", isPossible)
		fmt.Printf("\n")

	}

	// Return the total
	fmt.Printf("Sum of all IDs is %d\n", idTotal)
}

func getColourValue(subGame, colour string) int {
	value := 0

	// Get index of appearance of the colour
	colourIndex := strings.Index(subGame, colour)

	// Extract the value
	if colourIndex > 0 {
		colourStringValue := string(subGame[colourIndex-3]) + string(subGame[colourIndex-2])

		// Trim whitespace
		colourStringValue = strings.TrimSpace(colourStringValue)

		// Cast to int
		colourIntValue, err := strconv.Atoi(colourStringValue)
		if err != nil {
			log.Fatal(err)
		}

		value += colourIntValue
	}

	return value
}
