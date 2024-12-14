package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

func main()  {
	fileFlag := flag.String("input", "input.txt", "Name of the file containing the reports")
	program := getProgram(*fileFlag)
	fmt.Println(findMatches(program))
}

func getProgram(fileName string) string {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	var resultString string = ""
	for scanner.Scan() {
		resultString += scanner.Text()
	}
	return resultString
}

func findMatches(program string) int {
	var result int = 0
	var regex string = "mul\\((\\d{1,3}),(\\d{1,3})\\)|do\\(\\)|don't\\(\\)"
	var isDo bool = true
	programRegexp, _ := regexp.Compile(regex)
	for _, v := range programRegexp.FindAllStringSubmatch(program, -1) {
		if v[0] == "do()" {
			isDo = true
			continue
		} else if v[0] == "don't()" {
			isDo = false
			continue
		}
		if !isDo {
			continue
		}
		firstNumber, _ := strconv.Atoi(v[1])
		secondNumber, _ := strconv.Atoi(v[2])
		mul := firstNumber * secondNumber
		result += mul
		
	}
	return result
}
