package day1

import (
	"bufio"
	"container/heap"
	"log"
	"math"
	"strconv"
	"strings"

	"github.com/luisya22/aoc2023/datastruct"
	"github.com/luisya22/aoc2023/fileman"
	"github.com/spf13/cobra"
)

var aCmd = &cobra.Command{
	Use:   "a",
	Short: "Year 2024, Day 1, Problem A",
	Run: func(cmd *cobra.Command, args []string) {

		scanner, inputFile := fileman.GetFileLineBuffer("cmd/year2024day1/input.txt")
		defer inputFile.Close()

		result := partA(scanner)

		log.Println("Result A: ", result)
	},
}

func partA(inputScanner *bufio.Scanner) int {
	distanceSum := 0

	leftList := &datastruct.IntHeap{}
	heap.Init(leftList)

	rightList := &datastruct.IntHeap{}
	heap.Init(rightList)

	for inputScanner.Scan() {
		line := inputScanner.Text()

		splittedText := strings.Split(line, "   ")

		left, err := strconv.Atoi(splittedText[0])
		if err != nil {
			log.Fatalf("error converting %s to int", splittedText[0])
		}

		right, err := strconv.Atoi(splittedText[1])
		if err != nil {
			log.Fatalf("error converting %s to int", splittedText[1])
		}

		heap.Push(leftList, left)
		heap.Push(rightList, right)
	}

	length := leftList.Len()

	for x := 0; x < length; x++ {
		left := heap.Pop(leftList).(int)
		right := heap.Pop(rightList).(int)
		diff := float64(left - right)

		distanceSum += int(math.Abs(diff))
	}

	return distanceSum
}
