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

			red := getRed(subgame)
			if red > highestRed {
				highestRed = red
			}
			green := getGreen(subgame)
			if green > highestGreen {
				highestGreen = green
			}
			blue := getBlue(subgame)
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

func getRed(subGame string) int {
	total := 0

	// Get index of appearance of the word
	redWordIndex := strings.Index(subGame, "red")

	// Extract the value
	if redWordIndex > 0 {
		redWordValue := string(subGame[redWordIndex-3]) + string(subGame[redWordIndex-2])

		// Trim whitespace
		redWordValue = strings.TrimSpace(redWordValue)

		// Cast to int
		red, err := strconv.Atoi(redWordValue)
		if err != nil {
			log.Fatal(err)
		}

		total += red
	}

	return total
}

func getGreen(subGame string) int {
	total := 0

	// Get index of appearance of the word
	greenWordIndex := strings.Index(subGame, "green")

	// Extract the value
	if greenWordIndex > 0 {
		greenWordValue := string(subGame[greenWordIndex-3]) + string(subGame[greenWordIndex-2])

		// Trim whitespace
		greenWordValue = strings.TrimSpace(greenWordValue)

		// Cast to int
		green, err := strconv.Atoi(greenWordValue)
		if err != nil {
			log.Fatal(err)
		}

		total += green
	}

	return total
}

func getBlue(subGame string) int {
	total := 0

	// Get index of appearance of the word
	blueWordIndex := strings.Index(subGame, "blue")

	// Extract the value
	if blueWordIndex > 0 {
		blueWordValue := string(subGame[blueWordIndex-3]) + string(subGame[blueWordIndex-2])

		// Trim whitespace
		blueWordValue = strings.TrimSpace(blueWordValue)

		// Cast to int
		blue, err := strconv.Atoi(blueWordValue)
		if err != nil {
			log.Fatal(err)
		}

		total += blue
	}

	return total
}
