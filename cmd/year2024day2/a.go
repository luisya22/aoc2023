package day2

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"

	"github.com/luisya22/aoc2023/fileman"
	"github.com/spf13/cobra"
)

var aCmd = &cobra.Command{
	Use:   "a",
	Short: "Year 2024, Day 2, Problem A",
	Run: func(cmd *cobra.Command, args []string) {

		scanner, inputFile := fileman.GetFileLineBuffer("cmd/year2024day2/input.txt")
		defer inputFile.Close()

		result := partA(scanner)

		log.Println("Result A: ", result)
	},
}

func partA(inputScanner *bufio.Scanner) int {

	safeReports := 0

	for inputScanner.Scan() {
		line := inputScanner.Text()

		levels := strings.Split(line, " ")

		decreasing := false
		safeReport := true

		for x := 0; x < len(levels)-1; x++ {
			l1, err := strconv.Atoi(levels[x])
			if err != nil {
				log.Fatalf("error converting %s to int", levels[x])
			}

			l2, err := strconv.Atoi(levels[x+1])
			if err != nil {
				log.Fatalf("error converting %s to int", levels[x+1])
			}

			if x == 0 {
				// Set report direction
				if l1 > l2 {
					decreasing = true
				}
			}

			flDiff := float64(l1 - l2)
			diff := int(math.Abs(flDiff))

			case1 := decreasing && l2 > l1
			case2 := !decreasing && l1 > l2
			case3 := diff < 1 || diff > 3

			if case1 || case2 || case3 {
				safeReport = false
				break
			}
		}

		fmt.Println(line, safeReport)

		if safeReport {
			safeReports++
		}
	}

	return safeReports
}
