package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

// returns true if one pair is fully contained in the other and false otherwise
func part1(pair []string) bool {
	elf1 := strings.Split(pair[0], "-")
	elf2 := strings.Split(pair[1], "-")

	low1, _ := strconv.Atoi(elf1[0])
	high1, _ := strconv.Atoi(elf1[1])

	low2, _ := strconv.Atoi(elf2[0])
	high2, _ := strconv.Atoi(elf2[1])

	return (low1 <= low2 && high2 <= high1) || (low2 <= low1 && high1 <= high2)
}

func part2(pair []string) bool {
	elf1 := strings.Split(pair[0], "-")
	elf2 := strings.Split(pair[1], "-")

	low1, _ := strconv.Atoi(elf1[0])
	high1, _ := strconv.Atoi(elf1[1])

	low2, _ := strconv.Atoi(elf2[0])
	high2, _ := strconv.Atoi(elf2[1])

	return !((high1 < low2) || (high2 < low1))
}

func main() {
	inputFile := os.Args[1]
	file, err := ioutil.ReadFile(inputFile)

	if err != nil {
		fmt.Printf("Error reading file: %s\n", err)
	}

	// convert the file binary into a string using string
	fileContent := string(file)

	elfPairs := strings.Split(fileContent, "\n")
	numPairs := 0

	for i := range elfPairs {
		if part1(strings.Split(elfPairs[i], ",")) {
			numPairs += 1
		}
	}
	fmt.Println(numPairs)

	numPairs = 0
	for i := range elfPairs {
		if part2(strings.Split(elfPairs[i], ",")) {
			numPairs += 1
		}
	}
	fmt.Println(numPairs)
}
