package day4

import (
	"log"

	"github.com/luisya22/aoc2023/fileman"
	"github.com/spf13/cobra"
)

var XMAS = []string{"X", "M", "A", "S"}

var aCmd = &cobra.Command{
	Use:   "a",
	Short: "Year 2024, Day 4, Problem A",
	Run: func(cmd *cobra.Command, args []string) {

		fileText := fileman.GetFileAsString("cmd/year2024day4/input.txt")

		result := partA(fileText)

		log.Println("Result A: ", result)
	},
}

func partA(fileText string) int {
	appear := 0

	var wordSlice [][]string
	var line []string

	for _, char := range fileText {
		strChar := string(char)
		if strChar == "\n" {
			tempLine := make([]string, len(line))
			copy(tempLine, line)

			wordSlice = append(wordSlice, tempLine)

			line = line[:0]

			continue
		}

		line = append(line, strChar)
	}

	for x := 0; x < len(wordSlice); x++ {
		for y := 0; y < len(wordSlice[x]); y++ {
			appear += isCorrectWord(wordSlice, x, y, 0, -1, -1)
			appear += isCorrectWord(wordSlice, x, y, 0, -1, 0)
			appear += isCorrectWord(wordSlice, x, y, 0, -1, 1)
			appear += isCorrectWord(wordSlice, x, y, 0, 0, 1)
			appear += isCorrectWord(wordSlice, x, y, 0, 1, 1)
			appear += isCorrectWord(wordSlice, x, y, 0, 1, 0)
			appear += isCorrectWord(wordSlice, x, y, 0, 1, -1)
			appear += isCorrectWord(wordSlice, x, y, 0, 0, -1)

		}
	}

	return appear
}

func isCorrectWord(wordSlice [][]string, x int, y int, nextLetterIdx int, xNext int, yNext int) int {
	// TOP LEFT, TOP, TOP RIGHT
	// RIGHT
	// BOTTOM RIGHT, BOTTOM, BOTTOM LEFT
	// LEFT

	if nextLetterIdx == 4 {
		return 1
	}

	if x == -1 || y == -1 || x == len(wordSlice) || y == len(wordSlice[x]) {
		return 0
	}

	if wordSlice[x][y] != XMAS[nextLetterIdx] {
		return 0
	}

	nextLetterIdx++

	return isCorrectWord(wordSlice, x+xNext, y+yNext, nextLetterIdx, xNext, yNext)
}
