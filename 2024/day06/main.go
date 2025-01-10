package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"errors"
)

type GuardPosition struct {
	direction string
	posX      int
	posY      int
}

var maxNumberOfSteps int

func main() {
	fileFlag := flag.String("input", "input.txt", "Name of the file containing the text")
	// officeMap[y][x]
	var officeMap [][]string = readFile(*fileFlag)
	maxNumberOfSteps = len(officeMap) * len(officeMap[0])
	guardInfo, err := searchGuard(officeMap)
	if err != nil {
		fmt.Println(err)
	}
	numberOfPossibleLoops := calculatePossibleLoops(officeMap, guardInfo)
	officeMap,_ = letGuardRun(officeMap, guardInfo)
	count := countVisitedPositions(officeMap)
	fmt.Println(count)
	fmt.Println(numberOfPossibleLoops)
}

func readFile(fileName string) [][]string {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var officeMap [][]string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), "")
		officeMap = append(officeMap, line)
	}
	return officeMap
}

func calculatePossibleLoops(officeMap [][]string, guardInfo GuardPosition) int {
	var numberOfPossibleLoopMaps int = 0
	for y := 0; y < len(officeMap); y++ {
		for x := 0; x < len(officeMap[y]); x++ {
			if officeMap[y][x] == "^" || officeMap[y][x] == "#" {
				continue
			}
			copyOfficeMap := copyOfficeMap(officeMap)
			copyOfficeMap[y][x] = "#"
			_, err := letGuardRun(copyOfficeMap, guardInfo)
			if err != nil {
				numberOfPossibleLoopMaps += 1
			}
		}
	}
	return numberOfPossibleLoopMaps
}

func searchGuard(officeMap [][]string) (GuardPosition, error) {
	var guardInfo GuardPosition
	for idxY, line := range officeMap {
		for idxX, symbol := range line {
			if symbol == "^" {
				guardInfo.direction = "UP"
				guardInfo.posX = idxX
				guardInfo.posY = idxY
				return guardInfo, nil
			}
		}
	}
	return GuardPosition{"UP", -1, -1}, errors.New("No guard found.")
}

func letGuardRun(officeMap [][]string, guardInfo GuardPosition) ([][]string, error) {
	for i := 0; i < maxNumberOfSteps; i++ {
		officeMap[guardInfo.posY][guardInfo.posX] = "X"
		if wouldRunOutOfMap(officeMap, guardInfo) {
			return officeMap, nil
		}
		for isBlocked(officeMap, guardInfo) {
			guardInfo = rotateGuard(guardInfo)
		}
		officeMap, guardInfo = moveGuard(officeMap, guardInfo)
	}
	return nil, errors.New("To much looping")
}

func wouldRunOutOfMap(officeMap [][]string, guardInfo GuardPosition) bool {
	if (guardInfo.posY == 0 && guardInfo.direction == "UP") ||
	(guardInfo.posY == len(officeMap)-1 && guardInfo.direction == "DOWN") ||
	(guardInfo.posX == 0 && guardInfo.direction == "LEFT") ||
	(guardInfo.posX == len(officeMap[guardInfo.posY])-1 && guardInfo.direction == "RIGHT") {
		return true
	}
	return false
}

func moveGuard(officeMap [][]string, guardInfo GuardPosition) ([][]string, GuardPosition) {
	newY, newX := nextPos(guardInfo)
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
	return officeMap, guardInfo
}

/**
* Returns y,x
*/
func nextPos(guardInfo GuardPosition) (int, int) {
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

func isBlocked(officeMap [][]string, guardInfo GuardPosition) bool {
	if y,x := nextPos(guardInfo); officeMap[y][x] == "#" {
		return true
	}
	return false
}

func rotateGuard(guardInfo GuardPosition) GuardPosition {
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
	return guardInfo
}

func countVisitedPositions(officeMap [][]string) int {
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

func copyOfficeMap(officeMap [][]string) [][]string {
	copyOM := make([][]string, len(officeMap))
	for i, inner := range officeMap {
        copyOM[i] = make([]string, len(inner))
        copy(copyOM[i], inner)
    }
	return copyOM
}
