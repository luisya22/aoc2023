package day3

import (
	"bufio"
	"fmt"
	"log"
	"strconv"
	"unicode/utf8"

	"github.com/luisya22/aoc2023/fileman"
	"github.com/spf13/cobra"
)

var aCmd = &cobra.Command{
	Use:   "a",
	Short: "Year 2024, Day 3, Problem A",
	Run: func(cmd *cobra.Command, args []string) {

		scanner, inputFile := fileman.GetFileRuneBuffer("cmd/year2024day3/input.txt")
		defer inputFile.Close()

		result := partA(scanner)

		log.Println("Result A: ", result)
	},
}

func partA(inputScanner *bufio.Scanner) int {

	programSum := 0
	command := ""

	for inputScanner.Scan() {
		char := inputScanner.Text()

		if utf8.RuneCountInString(command) < 3 {
			command += char
			continue
		}

		if utf8.RuneCountInString(command) == 3 && command == "mul" {
			res, lastChar, err := getEvaluation(inputScanner)
			if err != nil {
				fmt.Println(err)
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

func getEvaluation(inputScanner *bufio.Scanner) (int, string, error) {

	if inputScanner.Text() != "(" {
		return 0, inputScanner.Text(), fmt.Errorf("invalid evaluation: missing first parenthesis")
	}

	// Get first number
	firstNumber := ""
	for inputScanner.Scan() {
		char := inputScanner.Text()
		if char == "," {
			break
		}

		// Check is number
		_, err := strconv.Atoi(char)
		if err != nil {
			fmt.Println(err)
			return 0, char, fmt.Errorf("found non number")
		}

		firstNumber += inputScanner.Text()

		if utf8.RuneCountInString(firstNumber) == 3 {
			if inputScanner.Scan() && inputScanner.Text() != "," {
				return 0, "", fmt.Errorf("comma not found after three digits")
			}
			break
		}
	}

	num1, err := strconv.Atoi(firstNumber)
	if err != nil {
		return 0, "", fmt.Errorf("failed to convert first number")
	}

	// Get second number
	secondNumber := ""
	for inputScanner.Scan() {
		char := inputScanner.Text()
		if char == ")" {
			break
		}

		// Check is number
		_, err := strconv.Atoi(char)
		if err != nil {
			return 0, char, fmt.Errorf("found non number")
		}

		secondNumber += inputScanner.Text()

		if utf8.RuneCountInString(secondNumber) == 3 {

			if inputScanner.Scan() && inputScanner.Text() != ")" {
				return 0, "", fmt.Errorf("closing parenthesis not found after three digits")
			}
			break
		}
	}

	num2, err := strconv.Atoi(secondNumber)
	if err != nil {
		return 0, "", fmt.Errorf("failed to convert second number")
	}

	fmt.Println("Evaluating", num1, num2, num1*num2)

	return num1 * num2, "", nil
}
