package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
)

func main()  {
	fileFlag := flag.String("input", "input.txt", "Name of the file containing the text")
	// textField [idxCol][idxRow]
	var textField [][]string = readFile(*fileFlag)
	var count int = countString(textField)
	fmt.Println(count)
}

func readFile(fileName string) [][]string {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	var result [][]string
	for scanner.Scan() {
		line := scanner.Text()
		var lineSlice []string
		for _, c := range line {
			lineSlice = append(lineSlice, string(c))
		}
		result = append(result, lineSlice)
	}
	return result
}

func countString(textField [][]string) int {
	var count int = 0
	for i := 1; i < len(textField)-1; i++ {
		for j := 1; j < len(textField[i])-1; j++ {
			if checkForCrossString(textField, i, j) {
				count += 1
			}
		}
	}
	return count
}

func checkForCrossString(
	textField [][]string,
	idxCol int,
	idxRow int) bool {
	if textField[idxCol][idxRow] != "A" {
		return false
	}
	if isMAS(textField[idxCol-1][idxRow-1], textField[idxCol+1][idxRow+1]) &&
		isMAS(textField[idxCol-1][idxRow+1], textField[idxCol+1][idxRow-1]) {
		return true
	}
	return false
}

func isMAS(topCorner string, bottomCorner string) bool {
	if (topCorner == "M" && bottomCorner == "S") || (topCorner == "S" && bottomCorner == "M") {
		return true
	}
	return false
}
