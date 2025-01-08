package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

type GuardPosition struct {
	direction string
	posX      int
	posY      int
}

// officeMap[y][x]
var officeMap [][]string
var guardInfo GuardPosition

func main() {
	fileFlag := flag.String("input", "input.txt", "Name of the file containing the text")
	officeMap = readFile(*fileFlag)
	searchGuard()
	letGuardRun()
	count := countVisitedPositions()
	fmt.Println(count)
}

func readFile(fileName string) [][]string {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), "")
		officeMap = append(officeMap, line)
	}
	return officeMap
}

func searchGuard() {
	for idxY, line := range officeMap {
		for idxX, symbol := range line {
			if symbol == "^" {
				guardInfo.direction = "UP"
				guardInfo.posX = idxX
				guardInfo.posY = idxY
				return
			}
		}
	}
}

func letGuardRun() {
	for {
		officeMap[guardInfo.posY][guardInfo.posX] = "X"
		if wouldRunOutOfMap() {
			return
		}
		for isBlocked() {
			rotateGuard()
		}
		moveGuard()
	}
}

func wouldRunOutOfMap() bool {
	if (guardInfo.posY == 0 && guardInfo.direction == "UP") ||
	(guardInfo.posY == len(officeMap)-1 && guardInfo.direction == "DOWN") ||
	(guardInfo.posX == 0 && guardInfo.direction == "LEFT") ||
	(guardInfo.posX == len(officeMap[guardInfo.posY])-1 && guardInfo.direction == "RIGHT") {
		return true
	}
	return false
}

func moveGuard() {
	newY, newX := nextPos()
	var symbol string
	switch guardInfo.direction {
	case "UP":
		symbol = "^"
	case "DOWN":
		symbol = "v"
	case "LEFT":
		symbol = "<"
	case "RIGHT":
		symbol = ">"
	default:
		log.Fatal("Broken direction: ", guardInfo.direction)
		symbol = "^"
	}
	officeMap[newY][newX] = symbol
	guardInfo.posX = newX
	guardInfo.posY = newY
}

/**
* Returns y,x
*/
func nextPos() (int, int) {
	switch guardInfo.direction {
	case "UP":
		return guardInfo.posY-1,guardInfo.posX
	case "DOWN":
		return guardInfo.posY+1,guardInfo.posX
	case "LEFT":
		return guardInfo.posY,guardInfo.posX-1
	case "RIGHT":
		return guardInfo.posY,guardInfo.posX+1
	default:
		log.Fatal("Broken direction: ", guardInfo.direction)
		return -1, -1
	}
}

func isBlocked() bool {
	if y,x := nextPos(); officeMap[y][x] == "#" {
		return true
	}
	return false
}

func rotateGuard() {
	switch guardInfo.direction {
	case "UP":
		guardInfo.direction = "RIGHT"
	case "DOWN":
		guardInfo.direction = "LEFT"
	case "LEFT":
		guardInfo.direction = "UP"
	case "RIGHT":
		guardInfo.direction = "DOWN"
	default:
		log.Fatal("Broken direction: ", guardInfo.direction)
	}
}

func countVisitedPositions() int {
	count := 0
	for y, _ := range officeMap {
		for x, _ := range officeMap[y] {
			if officeMap[y][x] == "X" {
				count += 1
			}
		}
	}
	return count
}
