package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"sort"
)

func main() {
	fileFlag := flag.String("input", "input.txt", "Name of the input file containing two vertical lists")
	list1, list2 := readInLists(*fileFlag)

	sortIntList(list1)
	sortIntList(list2)
	resultingDistance := calculateDistance(list1, list2)
	fmt.Println(resultingDistance)
}

func readInLists(fileName string) ([]int, []int) {
	var list1, list2 []int

	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		var n1, n2 int
		_, err := fmt.Sscanf(line, "%d %d", &n1, &n2)
		if err != nil {
			log.Fatal(err)
		}
		list1 = append(list1, n1)
		list2 = append(list2, n2)
	}
	return list1, list2
}

func sortIntList(l []int) []int {
	sort.Slice(l, func(i, j int) bool {
		return l[i] < l[j]
	})
	return l
}

func calculateDistance(list1 []int, list2 []int) int {
	result := 0
	for i := 0; i < len(list1); i++ {
		result += distanceBetweenNumbers(list1[i], list2[i])
	}
	return result
}

func distanceBetweenNumbers(number1 int, number2 int) int {
	difference := number1 - number2
	if difference < 0 {
		difference *= -1
	}
	return difference
}
