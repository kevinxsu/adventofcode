package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strconv"
	"strings"
)

// strings mapped to int
func smapi(lst []string) []int {
	res := make([]int, len(lst))
	for i := range lst {
		res[i], _ = strconv.Atoi(lst[i])
	}
	return res
}

func main() {
	inputFile := os.Args[1]
	file, err := ioutil.ReadFile(inputFile)

	if err != nil {
		fmt.Printf("Error reading file: %s\n", err)
	}

	// convert the file binary into a string using string
	fileContent := string(file)

	// split the file content by each elf, represented with a blank line
	itemChunks := strings.Split(fileContent, "\n\n")

	// create a slice of calories corresponding to each elf
	elfCalories := make([]int, len(itemChunks))

	// keep a running max of calories
	// max := 0.0
	for i := range itemChunks {
		calories := smapi(strings.Split(itemChunks[i], "\n"))

		sum := 0
		for j := range calories {
			sum += calories[j]
		}
		elfCalories[i] = sum
		// max = math.Max(max, sum)
	}

	sort.Ints(elfCalories)
	numElves := len(elfCalories)
	fmt.Println("Part 1: ", elfCalories[len(elfCalories)-1])
	fmt.Println("Part 2: ", elfCalories[numElves-1]+elfCalories[numElves-2]+elfCalories[numElves-3])
	// fmt.Println(max)
}
