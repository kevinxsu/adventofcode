package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func part2(stream string) int {
	for i := 0; i < len(stream); i++ {
		m := make(map[string]int)
		for j := 0; j < 14; j++ {
			m[string(stream[i+j])] += 1
		}
		if len(m) == 14 {
			return i + 14
		}
	}
	return -1
}

func part1(stream string) int {
	for i := 0; i < len(stream); i++ {
		m := make(map[string]int)
		m[string(stream[i])] += 1
		m[string(stream[i+1])] += 1
		m[string(stream[i+2])] += 1
		m[string(stream[i+3])] += 1
		if len(m) == 4 {
			return i + 4
		}
	}
	return -1
}

func main() {
	inputFile := os.Args[1]
	file, err := ioutil.ReadFile(inputFile)

	if err != nil {
		fmt.Printf("Error reading file: %s\n", err)
	}

	// convert the file binary into a string using string
	fileContent := string(file)

	fmt.Println(part1(fileContent))
	fmt.Println(part2(fileContent))
}
