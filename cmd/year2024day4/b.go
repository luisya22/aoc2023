package day4

import (
	"log"

	"github.com/luisya22/aoc2023/fileman"
	"github.com/spf13/cobra"
)

var bCmd = &cobra.Command{
	Use:   "b",
	Short: "Year 2024, Day 4, Problem B",
	Run: func(cmd *cobra.Command, args []string) {

		fileText := fileman.GetFileAsString("cmd/year2024day4/input.txt")

		result := partB(fileText)

		log.Println("Result B: ", result)
	},
}

func partB(fileText string) int {
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
			if wordSlice[x][y] == "A" && isCorrectCross(wordSlice, x, y) {
				appear++
			}

		}
	}

	return appear
}

func isCorrectCross(wordSlice [][]string, x int, y int) bool {

	if x-1 == -1 || y-1 == -1 || x+1 == len(wordSlice) || y+1 == len(wordSlice[x]) {
		return false
	}

	tl := wordSlice[x-1][y-1]
	tr := wordSlice[x-1][y+1]
	br := wordSlice[x+1][y+1]
	bl := wordSlice[x+1][y-1]

	if tl == "M" && tr == "M" && br == "S" && bl == "S" {
		return true
	}

	if tl == "S" && tr == "M" && br == "M" && bl == "S" {
		return true
	}

	if tl == "M" && tr == "S" && br == "S" && bl == "M" {
		return true
	}

	if tl == "S" && tr == "S" && br == "M" && bl == "M" {
		return true
	}

	return false

}
