package day3

import (
	"bufio"
	"fmt"
	"log"
	"strconv"

	"github.com/luisya22/aoc2023/fileman"
	"github.com/spf13/cobra"
)

var bCmd = &cobra.Command{
	Use:   "b",
	Short: "Day 3, Problem B",
	Run: func(cmd *cobra.Command, args []string) {

		scanner, inputFile := fileman.GetFileLineBuffer("cmd/day3/input.txt")
		defer inputFile.Close()

		result := partB(scanner)

		log.Println("Result B: ", result)
	},
}

func partB(inputScanner *bufio.Scanner) int {
	gearRatios := 0

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

			isAdj, adjSymbols := isAdjacent(actualPos, engine.symbols)

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

				// Add partnumber
				for _, as := range adjSymbols {
					engine.addPart(partNum, as)
				}

			}
		}
	}

	for _, sy := range engine.symbols {
		if sy.kind == "*" && len(sy.parts) == 2 {
			ratio := sy.parts[0] * sy.parts[1]

			gearRatios += ratio
		}
	}

	return gearRatios
}
