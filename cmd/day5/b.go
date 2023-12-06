package day5

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

var bCmd = &cobra.Command{
	Use:   "b",
	Short: "Day 5, Problem B",
	Run: func(cmd *cobra.Command, args []string) {

		scanner, inputFile := fileman.GetFileLineBuffer("cmd/day5/input.txt")

		defer inputFile.Close()

		result := partB(scanner)

		log.Println("Result B: ", result)
	},
}

func partB(inputScanner *bufio.Scanner) uint64 {
	lineNum := 0
	seeds := []seedsType{}
	var err error
	seedMap := [][]seedMapLine{}

	actualMap := -1
	seedMap = append(seedMap, []seedMapLine{})
	for inputScanner.Scan() {

		line := inputScanner.Text()

		switch {
		case lineNum == 0:
			seeds, err = getSeedsRanges(line)
			if err != nil {
				fmt.Println(err.Error())
				return 0
			}
		case line == "":
			seedMap = append(seedMap, []seedMapLine{})
			actualMap++
		case isDigit(line[0]):
			seedMap[actualMap], err = addLine(seedMap[actualMap], line)
			if err != nil {
				fmt.Println(err.Error())
				return 0
			}
		default:

		}

		lineNum++
	}

	var smaller uint64 = math.MaxUint64

	for _, seed := range seeds {
		seedLocation := processSeeds(seed, seedMap)

		fmt.Println(seedLocation)
		if seedLocation < smaller {
			smaller = seedLocation
		}
	}

	return smaller
}

type seedsType struct {
	initialSeed uint64
	seedRange   uint64
}

func processSeeds(sm seedsType, seedMap [][]seedMapLine) uint64 {
	var min uint64 = math.MaxUint64

	for i := sm.initialSeed; i < sm.initialSeed+sm.seedRange; i++ {
		seedLocation := getSeedLocation(i, seedMap)

		if seedLocation < min {
			min = seedLocation
		}
	}

	return min
}

func getSeedLocation(seed uint64, seedMap [][]seedMapLine) uint64 {

	for _, seedMapLine := range seedMap {
		for _, line := range seedMapLine {
			if seed >= line.source && seed < line.source+line.mapRange {

				diff := seed - line.source

				newSeedNum := line.destination + diff

				seed = newSeedNum

				break
			}
		}

	}

	return seed
}

func getSeedsRanges(seedsStr string) ([]seedsType, error) {

	seedsArr := []seedsType{}

	seedsSplt := strings.Split(seedsStr, ":")

	seeds := strings.TrimSpace(seedsSplt[1])

	for i := 0; i < len(seeds); i++ {
		actualNum1 := ""
		actualNum2 := ""

		// Get First Number
		for i < len(seeds) && string(seeds[i]) != " " {
			actualNum1 += string(seeds[i])
			i++
		}

		num1, err := strconv.ParseUint(actualNum1, 10, 64)
		if err != nil {
			return []seedsType{}, err
		}

		i++

		// Get Second Number
		for i < len(seeds) && string(seeds[i]) != " " {
			actualNum2 += string(seeds[i])
			i++
		}

		num2, err := strconv.ParseUint(actualNum2, 10, 64)
		if err != nil {
			return []seedsType{}, err
		}

		sm := seedsType{
			initialSeed: num1,
			seedRange:   num2,
		}

		seedsArr = append(seedsArr, sm)
	}

	return seedsArr, nil
}
