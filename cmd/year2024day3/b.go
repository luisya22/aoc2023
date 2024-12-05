package day3

import (
	"bufio"
	"log"
	"strings"
	"unicode/utf8"

	"github.com/luisya22/aoc2023/fileman"
	"github.com/spf13/cobra"
)

var bCmd = &cobra.Command{
	Use:   "b",
	Short: "Year 2024, Day 3, Problem B",
	Run: func(cmd *cobra.Command, args []string) {

		scanner, inputFile := fileman.GetFileRuneBuffer("cmd/year2024day3/input.txt")
		defer inputFile.Close()

		result := partB(scanner)

		log.Println("Result B: ", result)
	},
}

func partB(inputScanner *bufio.Scanner) int {
	programSum := 0
	command := ""
	enabledCommand := ""
	enabled := true

	for inputScanner.Scan() {
		char := inputScanner.Text()

		if strings.Contains(enabledCommand, "do()") {
			enabled = true
			enabledCommand = ""
		}

		if strings.Contains(enabledCommand, "don't()") {
			enabled = false
			enabledCommand = ""
		}

		enabledCommand += char

		if utf8.RuneCountInString(command) < 3 {
			command += char
			continue
		}

		if utf8.RuneCountInString(command) == 3 && command == "mul" {

			if !enabled {
				command = ""
				continue
			}

			res, lastChar, err := getEvaluation(inputScanner)
			if err != nil {
				command = lastChar
			}

			programSum += res
		}

		if utf8.RuneCountInString(command) == 3 {
			command = command[1:]
			command += char
		}
	}

	return programSum
}

// getEvaluation is on part a
