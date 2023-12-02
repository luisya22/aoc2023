package day2

import (
	"bufio"
	"fmt"
	"log"

	"github.com/luisya22/aoc2023/fileman"
	"github.com/spf13/cobra"
)

var bCmd = &cobra.Command{
	Use:   "b",
	Short: "Day 2, Problem B",
	Run: func(cmd *cobra.Command, args []string) {

		scanner, inputFile := fileman.GetFileLineBuffer("cmd/day2/input.txt")
		defer inputFile.Close()

		result := partB(scanner)

		log.Println("Result B: ", result)
	},
}

func partB(inputScanner *bufio.Scanner) int {

	gamePowerSum := 0

	for inputScanner.Scan() {
		line := inputScanner.Text()

		game, err := parseGame(line)
		if err != nil {
			fmt.Println(err.Error())
			return -1
		}

		gamePower := 1

		minCubesToPlay := map[string]int{
			"red":   1,
			"green": 1,
			"blue":  1,
		}

		for _, gc := range game.gameCubes {
			for _, c := range gc.cubes {

				minCube := minCubesToPlay[c.color]
				if c.quantity > minCube {
					minCubesToPlay[c.color] = c.quantity
				}
			}
		}

		for _, minCubeQty := range minCubesToPlay {
			gamePower *= minCubeQty
		}

		gamePowerSum += gamePower

		fmt.Println(gamePowerSum, gamePower)
	}

	return gamePowerSum
}
