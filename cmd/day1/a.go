/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package day1

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/luisya22/aoc2023/fileman"
	"github.com/spf13/cobra"
)

// aCmd represents the a command
var aCmd = &cobra.Command{
	Use:   "a",
	Short: "Day 1, Problem A",
	Run: func(cmd *cobra.Command, args []string) {

		inputStr := fileman.GetFileAsString("cmd/day1/input.txt")
		splitStr := strings.Split(inputStr, "\n")

		result := partA(splitStr)

		log.Println("Result 1: ", result)
	},
}

func partA(splitStr []string) int {

	numbers := []int{}
	var sum int

	for _, line := range splitStr {
		firstDigit := "0"
		lastDigit := "0"

		// Find first digit
		for i := 0; i < len(line); i++ {
			if isDigit(line[i]) {
				firstDigit = string(line[i])
				break
			}
		}
		// Find last digit
		for i := len(line) - 1; i >= 0; i-- {
			if isDigit(line[i]) {
				lastDigit = string(line[i])
				break
			}
		}

		digit, err := strconv.Atoi(fmt.Sprintf("%s%s", firstDigit, lastDigit))
		if err != nil {
			return -1
		}

		numbers = append(numbers, digit)

	}

	for _, digit := range numbers {
		sum += digit
	}

	return sum
}
