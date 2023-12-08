package day6

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
	Short: "Day 6, Problem A",
	Run: func(cmd *cobra.Command, args []string) {

		scanner, inputFile := fileman.GetFileLineBuffer("cmd/day6/input.txt")
		defer inputFile.Close()

		result := partA(scanner)

		log.Println("Result A: ", result)
	},
}

func partA(inputScanner *bufio.Scanner) int {
	racesBeatMult := 1

	times, distances, err := parseRaces(inputScanner)
	if err != nil {
		fmt.Printf("error parsing: %v\n", err)
		return -1
	}

	for i := 0; i < len(times); i++ {
		t := times[i]
		d := distances[i]

		disBeatSum := 0
		for x := 0; x <= t; x++ {
			dis := x * (t - x)

			if dis > d {
				disBeatSum++
			}
		}

		racesBeatMult *= disBeatSum
	}

	return racesBeatMult
}

func parseRaces(inputScanner *bufio.Scanner) ([]int, []int, error) {
	times := []int{}
	distances := []int{}

	// Get the times
	inputScanner.Scan()
	timesLine := inputScanner.Text()

	timeStrSplt := strings.Split(timesLine, ":")

	timesStr := strings.TrimSpace(timeStrSplt[1])

	for i := 0; i < len(timesStr); i++ {
		if !isDigit(timesStr[i]) {
			continue
		}

		actualNum := ""
		for i < len(timesStr) && isDigit(timesStr[i]) {
			actualNum += string(timesStr[i])
			i++
		}

		timeNum, err := strconv.Atoi(actualNum)
		if err != nil {
			return []int{}, []int{}, fmt.Errorf("error parsing: %v; %v", actualNum, err)
		}

		times = append(times, timeNum)
	}

	// Get the times
	inputScanner.Scan()
	distancesLine := inputScanner.Text()

	distanceStrSplt := strings.Split(distancesLine, ":")

	distanceStr := strings.TrimSpace(distanceStrSplt[1])

	for i := 0; i < len(distanceStr); i++ {
		if !isDigit(distanceStr[i]) {
			continue
		}

		actualNum := ""
		for i < len(distanceStr) && isDigit(distanceStr[i]) {
			actualNum += string(distanceStr[i])
			i++
		}

		distNum, err := strconv.Atoi(actualNum)
		if err != nil {
			return []int{}, []int{}, fmt.Errorf("error parsing: %v; %v", actualNum, err)
		}

		distances = append(distances, distNum)
	}

	return times, distances, nil
}
