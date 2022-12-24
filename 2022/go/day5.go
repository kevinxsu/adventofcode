package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	inputFile := os.Args[1]
	file, err := ioutil.ReadFile(inputFile)

	if err != nil {
		fmt.Printf("Error reading file: %s\n", err)
	}

	// convert the file binary into a string using string
	fileContent := string(file)

	f := strings.Split(fileContent, "\n\n")
	// fmt.Println(f[0])

	layout := strings.Split(f[0], "\n")
	stackNumbers := strings.Split(strings.TrimSpace(layout[len(layout)-1]), " ")
	numStacks := stackNumbers[len(stackNumbers)-1]

	// fmt.Println(stackNumbers)
	fmt.Println(numStacks)

}
