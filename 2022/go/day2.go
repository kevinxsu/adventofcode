package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

// given opponent choice and your choice, calculate the score of the round
// func roundScore(opponent, you string) int {
func roundScore(round string) int {
	scoreMappings := make(map[string]int)
	scoreMappings["A X"] = 1 + 3
	scoreMappings["A Y"] = 2 + 6
	scoreMappings["A Z"] = 3 + 0
	scoreMappings["B X"] = 1 + 0
	scoreMappings["B Y"] = 2 + 3
	scoreMappings["B Z"] = 3 + 6
	scoreMappings["C X"] = 1 + 6
	scoreMappings["C Y"] = 2 + 0
	scoreMappings["C Z"] = 3 + 3
	return scoreMappings[round]
}

func roundScore2(round string) int {
	optLost := make(map[string]int)
	// rock
	// need to lose -> scissors
	optLost["A X"] = 3 + 0
	// need to tie
	optLost["A Y"] = 1 + 3
	// need to win
	optLost["A Z"] = 2 + 6

	// paper
	// need to lose -> rock
	optLost["B X"] = 1 + 0
	// need to tie -> paper
	optLost["B Y"] = 2 + 3
	// need to win -> scissors
	optLost["B Z"] = 3 + 6

	// scissors
	// need to lose -> paper
	optLost["C X"] = 2 + 0
	// need to tie -> scissors
	optLost["C Y"] = 3 + 3
	// need to win -> rock
	optLost["C Z"] = 1 + 6
	return optLost[round]
}

func main() {
	inputFile := os.Args[1]
	file, err := ioutil.ReadFile(inputFile)

	if err != nil {
		fmt.Printf("Error reading file: %s\n", err)
	}

	// convert the file binary into a string using string
	fileContent := string(file)

	// each round separated by a new line
	rounds := strings.Split(fileContent, "\n")

	score1 := 0
	score2 := 0
	for i := range rounds {
		score1 += roundScore(rounds[i])
		score2 += roundScore2(rounds[i])
	}

	fmt.Println(score1)
	fmt.Println(score2)

}
