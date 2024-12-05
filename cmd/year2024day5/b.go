package day5

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
	Short: "Year 2024, Day 5, Problem B",
	Run: func(cmd *cobra.Command, args []string) {

		scanner, inputFile := fileman.GetFileLineBuffer("cmd/year2024day5/input.txt")
		defer inputFile.Close()

		result := partB(scanner)

		log.Println("Result B: ", result)
	},
}

type updateValue struct {
	realValue string
	lineValue int
}

func partB(inputScanner *bufio.Scanner) int {
	pageOrderingResults := 0

	instructionMap := make(map[string][]string)

	instructions := true
	for inputScanner.Scan() {
		line := inputScanner.Text()

		if line == "" {
			instructions = false
			continue
		}

		if instructions {
			splitLine := strings.Split(line, "|")

			if _, ok := instructionMap[splitLine[0]]; !ok {
				instructionMap[splitLine[0]] = []string{}
			}

			instructionMap[splitLine[0]] = append(instructionMap[splitLine[0]], splitLine[1])
		}

		if !instructions {
			splitLine := strings.Split(line, ",")

			appeared := make(map[string]interface{})

			rightOrder := true

			for _, d := range splitLine {
				for _, insA := range instructionMap[d] {
					if appeared[insA] != nil {
						rightOrder = false
					}
				}

				appeared[d] = true
			}

			if !rightOrder {
				fmt.Print(splitLine)
				pageOrderingResults += getOrderedMiddle(splitLine, instructionMap)
			}
		}

	}

	return pageOrderingResults
}

func getOrderedMiddle(splitLine []string, instructionMap map[string][]string) int {
	orderedLine := []string{}

	for _, d := range splitLine {
		if len(orderedLine) == 0 {
			orderedLine = append(orderedLine, d)
			continue
		}

		inserted := false
		for i, ol := range orderedLine {
			if !goAfter(d, instructionMap[ol]) {
				orderedLine = InsertBefore(orderedLine, d, i)
				inserted = true
				break
			}
		}

		if !inserted {
			orderedLine = append(orderedLine, d)
		}
	}

	middle, err := strconv.Atoi(orderedLine[len(orderedLine)/2])
	if err != nil {
		log.Fatalf("Can't convert %v to int", orderedLine[len(orderedLine)/2])
	}

	return middle
}

func goAfter(d string, instructions []string) bool {
	for _, insV := range instructions {
		if d == insV {
			return true
		}
	}

	return false
}

func InsertBefore(slice []string, element string, index int) []string {
	if index < 0 || index > len(slice) {
		panic("index out of range")
	}

	newSlice := make([]string, len(slice)+1)
	copy(newSlice[:index], slice[:index])
	newSlice[index] = element
	copy(newSlice[index+1:], slice[index:])
	return newSlice
}
