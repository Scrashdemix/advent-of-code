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

type Equation struct {
	result int
	terms  []int
}

var operators []string = []string{"+", "*", "|"}

func main() {
	fileFlag := flag.String("input", "input.txt", "Name of the file containing the text")
	var eqs []Equation = readFile(*fileFlag)

	println(generateResults(eqs))
}

func readFile(fileName string) []Equation {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var equations []Equation
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var current Equation
		line := strings.Split(scanner.Text(), " ")
		// equation result
		result, _ := strings.CutSuffix(line[0], ":")
		current.result, _ = strconv.Atoi(result)
		// equation terms
		var terms []int = make([]int, len(line)-1)
		for i := 0; i < len(line)-1; i++ {
			terms[i], _ = strconv.Atoi(line[i+1])
		}
		current.terms = terms
		equations = append(equations, current)
	}
	return equations
}

func generateCombinations(n int, current string, results *[]string) {
	if len(current) == n {
		*results = append(*results, current)
		return
	}
	for _, op := range operators {
		generateCombinations(n, current+op, results)
	}
}

func generateResults(equations []Equation) int {
	overallSum := 0
	numOfEqs := len(equations)
	for idx, eq := range equations {
		overallSum += checkEquation(eq)
		if idx%10 == 0 {
			fmt.Println("Progress: ", idx, "/", numOfEqs)
		}
	}
	return overallSum
}

func checkEquation(eq Equation) int {
	var operators []string
	generateCombinations(len(eq.terms)-1, "", &operators)
	for _, eqOperators := range operators {
		res := calcFullEquation(eq.terms, eqOperators)
		if res == eq.result {
			return res
		}
	}
	return 0
}

func calcFullEquation(terms []int, operators string) int {
	var currentResult int = terms[0]
	for idx, op := range operators {
		if op == '+' {
			currentResult += terms[idx+1]
		} else if op == '*' {
			currentResult *= terms[idx+1]
		} else if op == '|' {
			temp := strconv.Itoa(currentResult) + strconv.Itoa(terms[idx+1])
			currentResult, _ = strconv.Atoi(temp)
		}
	}
	return currentResult
}
