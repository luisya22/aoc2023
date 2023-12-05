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

var bCmd = &cobra.Command{
	Use:   "b",
	Short: "Day 4, Problem B",
	Run: func(cmd *cobra.Command, args []string) {

		scanner, inputFile := fileman.GetFileLineBuffer("cmd/day4/input.txt")
		defer inputFile.Close()

		result := partB(scanner)

		log.Println("Result B: ", result)
	},
}

func partB(inputScanner *bufio.Scanner) int {
	cardInstancesSum := 0

	cardInstances := make(map[int]int)
	actualCard := 0
	for inputScanner.Scan() {
		actualCard++
		card := inputScanner.Text()

		cardInstances[actualCard]++

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
		numbersWon := 0
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

				numbersWon += 1
			}
		}

		copies := cardInstances[actualCard]

		for i := copies; i > 0; i-- {
			sumCard := actualCard + 1
			for remainingWon := numbersWon; remainingWon > 0; remainingWon-- {
				cardInstances[sumCard]++
				sumCard++
			}
		}
	}

	for _, num := range cardInstances {
		cardInstancesSum += num
	}

	return cardInstancesSum
}
