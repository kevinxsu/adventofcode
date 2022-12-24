package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

// this assumes that there is only one item that appears in both compartments
func part1(rucksackItems string) int {
	// create the priority mapping
	priorityMapping := make(map[string]int)
	for i := 1; i <= 26; i++ {
		priorityMapping[string(97+i-1)] = i
		// priorityMapping[rune(97+i-1)] = i
	}
	for i := 27; i <= 52; i++ {
		priorityMapping[strings.ToUpper(string(97+i-27))] = i
		// priorityMapping[rune(97+i-27)] = i
	}
	// fmt.Println(priorityMapping)
	items := strings.Split(rucksackItems, "")
	compartmentSize := len(items) / 2

	first := rucksackItems[0:compartmentSize]
	second := rucksackItems[compartmentSize:]

	// O(n^2) but its probably fine
	for i := range first {
		for j := range second {
			if first[i] == second[j] {
				return priorityMapping[string(first[i])]
			}
		}
	}

	//fmt.Println(compartmentSize)
	//fmt.Println(items)
	//fmt.Println(first)
	//fmt.Println(second)
	//itemsMapping := make(map[string]int)
	//for i := 0; i < compartmentSize; i++ {
	//	itemsMapping[first[i]] += 1

	//}
	return 0
}

func part2(sack1, sack2, sack3 string) int {
	priorityMapping := make(map[string]int)
	for i := 1; i <= 26; i++ {
		priorityMapping[string(97+i-1)] = i
		// priorityMapping[rune(97+i-1)] = i
	}
	for i := 27; i <= 52; i++ {
		priorityMapping[strings.ToUpper(string(97+i-27))] = i
		// priorityMapping[rune(97+i-27)] = i
	}
	for i := range sack1 {
		for j := range sack2 {
			for k := range sack3 {
				if sack1[i] == sack2[j] && sack2[j] == sack3[k] {
					return priorityMapping[string(sack1[i])]
				}
			}
		}
	}
	return 0
}

func main() {
	inputFile := os.Args[1]
	file, err := ioutil.ReadFile(inputFile)

	if err != nil {
		fmt.Printf("Error reading file: %s\n", err)
	}

	// convert the file binary into a string using string
	fileContent := string(file)
	rucksacks := strings.Split(fileContent, "\n")
	priorities := 0
	for i := range rucksacks {
		// fmt.Println(rucksacks[i])
		priorities += part1(rucksacks[i])
	}
	fmt.Println(priorities)

	priorities = 0
	for i := 0; i < len(rucksacks); i += 3 {
		priorities += part2(rucksacks[i], rucksacks[i+1], rucksacks[i+2])
	}
	fmt.Println(priorities)
}
