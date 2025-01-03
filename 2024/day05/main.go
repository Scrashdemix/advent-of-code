package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

var rules map [int][]int

func main()  {
	fileFlag := flag.String("input", "input.txt", "Name of the file containing the text")
	updates := readFile(*fileFlag)
	var result = checkUpdates(updates)
	fmt.Println(result)
}

func readFile(fileName string) [][]int {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	parseRulesFromFile(scanner)
	updates := parseUpdatesFromFile(scanner)
	return updates
}

func parseRulesFromFile (scanner *bufio.Scanner) {
	pagesBefore := make(map[int][]int)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}
		ruleNumbers := strings.Split(line, "|")
		firstNumber, _ := strconv.Atoi(ruleNumbers[0])
		secondNumber, _ := strconv.Atoi(ruleNumbers[1])
		pagesBefore[secondNumber] = append(pagesBefore[secondNumber], firstNumber)
	}
	rules = pagesBefore
}

func parseUpdatesFromFile (scanner *bufio.Scanner) [][]int {
	var result [][]int = make([][]int, 0)
	for scanner.Scan() {
		line := scanner.Text()
		pages := strings.Split(line, ",")
		result = append(result, convToIntSlice(pages))
	}
	return result
}

func convToIntSlice(str []string) []int {
	var result []int = make([]int, len(str))
	for idx, s := range str {
		result[idx], _ = strconv.Atoi(s)
	}
	return result
}

func checkUpdates(updates [][]int) int {
	var result int = 0
	for _, update := range updates {
		result += checkSingleUpdate(update)
	}
	return result
}

func checkSingleUpdate(update []int) int {
	for idx, _ := range update {
		if !isInLineWithRules(update, idx) {
			return 0
		}
	}
	return update[len(update)/2]
}

func isInLineWithRules(update []int, idx int) bool {
	number := update[idx]
	for _, n := range rules[number] {
		if slices.Index(update, n) > idx {
			return false
		}
	}
	return true
}
