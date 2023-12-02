package day2

import (
	"bufio"
	"fmt"
	"log"

	"github.com/luisya22/aoc2023/fileman"
	"github.com/spf13/cobra"
)

var aCmd = &cobra.Command{
	Use:   "a",
	Short: "Day 2, Problem A",
	Run: func(cmd *cobra.Command, args []string) {

		scanner, inputFile := fileman.GetFileLineBuffer("cmd/day2/input.txt")
		defer inputFile.Close()

		result := partA(scanner)

		log.Println("Result A: ", result)
	},
}

func partA(inputScanner *bufio.Scanner) int {

	maxCubes := map[string]int{
		"red":   12,
		"green": 13,
		"blue":  14,
	}

	possibleGamesIdSum := 0

	for inputScanner.Scan() {
		line := inputScanner.Text()

		possible := true

		game, err := parseGame(line)
		if err != nil {
			fmt.Println(err.Error())
			return -1
		}

	out:
		for _, gc := range game.gameCubes {
			for _, c := range gc.cubes {

				max := maxCubes[c.color]
				if c.quantity > max {
					possible = false
					break out
				}
			}
		}

		if possible {
			possibleGamesIdSum += game.id
		}
	}

	return possibleGamesIdSum
}
