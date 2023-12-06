package day5

import (
	"bufio"
	"fmt"
	"log"

	"github.com/luisya22/aoc2023/fileman"
	"github.com/spf13/cobra"
)

var aCmd = &cobra.Command{
	Use:   "a",
	Short: "Day 5, Problem A",
	Run: func(cmd *cobra.Command, args []string) {

		scanner, inputFile := fileman.GetFileLineBuffer("cmd/day5/input.txt")
		defer inputFile.Close()

		result := partA(scanner)

		log.Println("Result A: ", result)
	},
}

func partA(inputScanner *bufio.Scanner) uint64 {

	lineNum := 0
	seeds := []uint64{}
	var err error
	seedMap := []seedMapLine{}

	for inputScanner.Scan() {

		line := inputScanner.Text()

		switch {
		case lineNum == 0:
			seeds, err = getSeeds(line)
			if err != nil {
				fmt.Println(err.Error())
				return 0
			}
		case line == "":
			seeds, err = transferSeeds(seedMap, seeds)
			if err != nil {
				fmt.Println(err.Error())
				return 0
			}
			seedMap = []seedMapLine{}
		case isDigit(line[0]):
			seedMap, err = addLine(seedMap, line)
			if err != nil {
				fmt.Println(err.Error())
				return 0
			}
		default:

		}

		lineNum++
	}

	// After last scan transferSeeds
	seeds, err = transferSeeds(seedMap, seeds)
	if err != nil {
		fmt.Println(err.Error())
		return 0
	}

	smaller := seeds[0]
	for _, soil := range seeds {
		if soil < smaller {
			smaller = soil
		}
	}

	return smaller
}
