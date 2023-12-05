package day4

import (
	"bufio"
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/luisya22/aoc2023/fileman"
	"github.com/spf13/cobra"
)

var aCmd = &cobra.Command{
	Use:   "a",
	Short: "Day 4, Problem A",
	Run: func(cmd *cobra.Command, args []string) {

		scanner, inputFile := fileman.GetFileLineBuffer("cmd/day4/input.txt")
		defer inputFile.Close()

		result := partA(scanner)

		log.Println("Result A: ", result)
	},
}

func partA(inputScanner *bufio.Scanner) int {

	cardPointsSum := 0
	for inputScanner.Scan() {
		card := inputScanner.Text()

		// Split winning numbers from card numbers
		cardParts := strings.Split(card, ":")
		cardNums := strings.Split(cardParts[1], "|")

		winningNumbers := make(map[int]struct{})

		winningNumbersParts := strings.Split(strings.TrimSpace(cardNums[0]), " ")

		for _, wn := range winningNumbersParts {

			if wn == "" {
				continue
			}

			winningNumber, err := strconv.Atoi(wn)
			if err != nil {
				fmt.Println(err.Error())
				return -1
			}

			winningNumbers[winningNumber] = struct{}{}

		}

		cardNumbersParts := strings.Split(strings.TrimSpace(cardNums[1]), " ")
		cardPoints := 0
		for _, cn := range cardNumbersParts {

			if cn == "" {
				continue
			}

			cardNumber, err := strconv.Atoi(cn)
			if err != nil {
				fmt.Println(err.Error())
				return -1
			}

			if _, ok := winningNumbers[cardNumber]; ok {

				if cardPoints == 0 {
					cardPoints = 1
				} else {
					cardPoints *= 2
				}
			}
		}

		cardPointsSum += cardPoints
	}

	return cardPointsSum
}
