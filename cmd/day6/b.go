package day6

import (
	"bufio"
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/luisya22/aoc2023/fileman"
	"github.com/spf13/cobra"
)

var bCmd = &cobra.Command{
	Use:   "b",
	Short: "Day 6, Problem B",
	Run: func(cmd *cobra.Command, args []string) {

		scanner, inputFile := fileman.GetFileLineBuffer("cmd/day6/input.txt")
		defer inputFile.Close()

		result := partB(scanner)

		log.Println("Result B: ", result)
	},
}

func partB(inputScanner *bufio.Scanner) int {
	racesBeatSum := 0

	timeNum, distanceNum, err := parseRace(inputScanner)
	if err != nil {
		fmt.Printf("error parsing: %v\n", err)
		return -1
	}

	for x := 0; x <= timeNum; x++ {
		dis := x * (timeNum - x)

		if dis > distanceNum {
			racesBeatSum++
		}
	}

	return racesBeatSum
}

func parseRace(inputScanner *bufio.Scanner) (int, int, error) {
	// Get the time
	inputScanner.Scan()
	timesLine := inputScanner.Text()

	timeStrSplt := strings.Split(timesLine, ":")

	timesStr := strings.TrimSpace(timeStrSplt[1])

	actualNum := ""
	for i := 0; i < len(timesStr); i++ {
		if isDigit(timesStr[i]) {
			actualNum += string(timesStr[i])
		}
	}

	timeNum, err := strconv.Atoi(actualNum)
	if err != nil {
		return -1, -1, fmt.Errorf("error parsing: %v; %v", actualNum, err)
	}

	// Get the distance
	inputScanner.Scan()
	distancesLine := inputScanner.Text()

	distanceStrSplt := strings.Split(distancesLine, ":")

	distanceStr := strings.TrimSpace(distanceStrSplt[1])

	actualNum = ""
	for i := 0; i < len(distanceStr); i++ {
		if isDigit(distanceStr[i]) {
			actualNum += string(distanceStr[i])
		}
	}

	distNum, err := strconv.Atoi(actualNum)
	if err != nil {
		return -1, -1, fmt.Errorf("error parsing: %v; %v", actualNum, err)
	}

	return timeNum, distNum, nil
}
