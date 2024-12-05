package day5

import (
	"bufio"
	"log"
	"strconv"
	"strings"

	"github.com/luisya22/aoc2023/fileman"
	"github.com/spf13/cobra"
)

var aCmd = &cobra.Command{
	Use:   "a",
	Short: "Year 2024, Day 5, Problem A",
	Run: func(cmd *cobra.Command, args []string) {

		scanner, inputFile := fileman.GetFileLineBuffer("cmd/year2024day5/input.txt")
		defer inputFile.Close()

		result := partA(scanner)

		log.Println("Result A: ", result)
	},
}

func partA(inputScanner *bufio.Scanner) int {
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

			if rightOrder {
				middle, err := strconv.Atoi(splitLine[len(splitLine)/2])
				if err != nil {
					log.Fatalf("non number: line -> %v, value: %v", line, splitLine[len(splitLine)/2])
				}

				pageOrderingResults += middle
			}
		}

	}

	return pageOrderingResults
}
