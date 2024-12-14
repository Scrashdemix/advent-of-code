package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main () {
	fileFlag := flag.String("input", "input.txt", "Name of the file containing the reports")
	reports := readInReports(*fileFlag)
	numberOfSafeReports := 0
	for _, r := range reports {
		if reportIsSafe(r) {
			numberOfSafeReports += 1
		} else {
			if secondChance(r) {
				numberOfSafeReports += 1
			}
		}
	}
	fmt.Println("Number of safe reports: ", numberOfSafeReports)
}

func readInReports(fileName string) [][]int {
	var result [][]int
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		result = append(result, parseLineToReport(line))
	}
	return result
}

func parseLineToReport(line string) []int {
	var rep []int
	for _, level := range strings.Fields(line) {
		l, err := strconv.Atoi(level)
		if err != nil {
			log.Fatal(err)
		}
		rep = append(rep, l)
	}
	return rep
}

func reportIsSafe(report []int) bool {
	var isIncreasing bool
	for idx, _ := range report {
		// Check whether the levels are still increasing/decreasing
		if idx == 0 {
			continue
		} else if idx == 1 {
			isIncreasing = report[idx-1] < report[idx]
		}
		if (isIncreasing && report[idx-1] > report[idx]) ||
		(!isIncreasing && report[idx-1] < report[idx]){
			return false
		}
		// Check whether difference is safe
		var diff int
		if isIncreasing {
			diff = report[idx] - report[idx-1]
		} else {
			diff = report[idx-1] - report[idx]
		}

		if diff < 1 || diff > 3 {
			return false
		}
	}
	return true
}

func secondChance(report []int) bool  {
	for i := 0; i < len(report); i++ {
		var tempArray []int = make([]int, len(report))
		copy(tempArray, report)
		tempArray = append(tempArray[:i], tempArray[i+1:]...)
		if reportIsSafe(tempArray) {
			return true
		}
	}
	return false
}
