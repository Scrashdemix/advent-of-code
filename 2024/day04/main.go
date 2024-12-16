package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
)

var directions [][]int = [][]int {
	{1, 0}, // down
	{0, 1}, // right
	{-1, 0}, // up
	{0, -1}, // left
	{1, 1}, // down-right
	{-1, -1}, // up-left
	{-1, 1}, // up-right
	{1, -1}, // down-left
}

func main()  {
	fileFlag := flag.String("input", "input.txt", "Name of the file containing the text")
	// textField [idxCol][idxRow]
	var textField [][]string = readFile(*fileFlag)
	var count int = countString(textField, "XMAS")
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

func countString(textField [][]string, searchString string) int {
	var count int = 0
	for i := 0; i < len(textField); i++ {
		for j := 0; j < len(textField[i]); j++ {
			for _, dir := range directions {
				resultingString := getStringFromField(textField, i, j, dir, len(searchString))
				if resultingString == searchString {
					count += 1
				}
			}
		}
	}
	return count
}

func getStringFromField(
	textField [][]string,
	idxCol int,
	idxRow int,
	direction []int,
	length int) string {
		if (direction[0] == -1 && idxCol < length-1) ||
		(direction[0] == 1 && (len(textField)-idxCol) <= length-1) ||
		(direction[1] == -1 && idxRow < length-1) ||
		(direction[1] == 1 && (len(textField[idxCol])-idxRow) <= length-1) {
			return ""
		}

		var tempString = ""
		for i := 0; i < length; i++ {
			tempString += textField[idxCol + i*direction[0]][idxRow + i*direction[1]]
		}
		return tempString
}
