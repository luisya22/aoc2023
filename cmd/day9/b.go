package day9

import (
	"bufio"
	"fmt"
	"log"

	"github.com/luisya22/aoc2023/fileman"
	"github.com/spf13/cobra"
)

var bCmd = &cobra.Command{
	Use:   "b",
	Short: "Day 9, Problem B",
	Run: func(cmd *cobra.Command, args []string) {

		scanner, inputFile := fileman.GetFileLineBuffer("cmd/day9/input.txt")
		defer inputFile.Close()

		result := partB(scanner)

		log.Println("Result B: ", result)
	},
}

func partB(inputScanner *bufio.Scanner) int {

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
			actualLine := predictionsData[i]
			if i == len(predictionsData)-1 {
				predictionsData[i] = append(predictionsData[i], 0)
				continue
			}

			num := actualLine[0] - predictionsData[i+1][0]
			actualLine = append([]int{num}, actualLine...)
			predictionsData[i] = actualLine
		}

		predictionsSum += predictionsData[0][0]
	}

	return predictionsSum
}
