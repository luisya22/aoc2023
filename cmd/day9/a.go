package day9

import (
	"bufio"
	"fmt"
	"log"

	"github.com/luisya22/aoc2023/fileman"
	"github.com/spf13/cobra"
)

var aCmd = &cobra.Command{
	Use:   "a",
	Short: "Day 9, Problem A",
	Run: func(cmd *cobra.Command, args []string) {

		scanner, inputFile := fileman.GetFileLineBuffer("cmd/day9/input.txt")
		defer inputFile.Close()

		result := partA(scanner)

		log.Println("Result A: ", result)
	},
}

func partA(inputScanner *bufio.Scanner) int {

	predictionsSum := 0
	for inputScanner.Scan() {
		line := inputScanner.Text()

		predictionsData, err := parseData(line)
		if err != nil {
			fmt.Println(err)
			return -1
		}

		nonZero := true
		counter := 0
		for nonZero {

			nonZero = false
			newLine := []int{}
			for i := 0; i < len(predictionsData[counter])-1; i++ {
				num := predictionsData[counter][i+1] - predictionsData[counter][i]

				newLine = append(newLine, num)

				if num != 0 {
					nonZero = true
				}
			}

			predictionsData = append(predictionsData, newLine)

			counter++
		}

		for i := len(predictionsData) - 1; i >= 0; i-- {
			if i == len(predictionsData)-1 {
				predictionsData[i] = append(predictionsData[i], 0)
				continue
			}

			num := predictionsData[i][len(predictionsData[i])-1] + predictionsData[i+1][len(predictionsData[i+1])-1]
			predictionsData[i] = append(predictionsData[i], num)
		}

		predictionsSum += predictionsData[0][len(predictionsData[0])-1]
	}

	return predictionsSum
}
