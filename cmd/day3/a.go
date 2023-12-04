package day3

import (
	"bufio"
	"fmt"
	"log"
	"strconv"

	"github.com/luisya22/aoc2023/fileman"
	"github.com/spf13/cobra"
)

var aCmd = &cobra.Command{
	Use:   "a",
	Short: "Day 3, Problem A",
	Run: func(cmd *cobra.Command, args []string) {

		scanner, inputFile := fileman.GetFileLineBuffer("cmd/day3/input.txt")
		defer inputFile.Close()

		result := partA(scanner)

		log.Println("Result A: ", result)
	},
}

func partA(inputScanner *bufio.Scanner) int {

	partNumSum := 0

	engine, err := parseEngine(inputScanner)
	if err != nil {
		fmt.Println(err.Error())
		return -1
	}

	for lineNumber, line := range engine.lines {
		for i := 0; i < len(line); i++ {
			actualPos := position{
				y: lineNumber,
				x: i,
			}

			isAdj, _ := isAdjacent(actualPos, engine.symbols)

			if isDigit(line[i]) && isAdj {
				// Go back to get beggining of number
				part := string(line[i])
				pointer := i - 1

				for pointer >= 0 && isDigit(line[pointer]) {
					part = fmt.Sprintf("%v%v", string(line[pointer]), part)
					pointer--
				}

				// Go front until end of number
				i++
				for i != len(line) && isDigit(line[i]) {
					part = fmt.Sprintf("%v%v", part, string(line[i]))
					i++
				}

				partNum, err := strconv.Atoi(part)
				if err != nil {
					fmt.Println(err.Error())
					return -1
				}

				partNumSum += partNum
			}
		}
	}

	return partNumSum
}
