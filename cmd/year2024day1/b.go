package day1

import (
	"bufio"
	"container/heap"
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/luisya22/aoc2023/datastruct"
	"github.com/luisya22/aoc2023/fileman"
	"github.com/spf13/cobra"
)

var bCmd = &cobra.Command{
	Use:   "b",
	Short: "Year 2024, Day 1, Problem B",
	Run: func(cmd *cobra.Command, args []string) {

		scanner, inputFile := fileman.GetFileLineBuffer("cmd/year2024day1/input.txt")
		defer inputFile.Close()

		result := partB(scanner)

		log.Println("Result B: ", result)
	},
}

func partB(inputScanner *bufio.Scanner) int {
	appereancesSum := 0

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

	appereancesMap := make(map[int]int, leftList.Len())
	left := heap.Pop(leftList).(int)
	right := heap.Pop(rightList).(int)
	appereances := 0
	isSameNumber := false

	for leftList.Len()+1 > 0 && rightList.Len() > 0 {
		_, f := appereancesMap[left]
		fmt.Println("is new left", left, isSameNumber, f)
		if amount, found := appereancesMap[left]; found && isSameNumber {
			appereancesSum += amount
			oldLeft := left
			left = heap.Pop(leftList).(int)
			isSameNumber = oldLeft == left
			continue
		}

		if left == right {

			appereances++
			right = heap.Pop(rightList).(int)
		} else {

			fmt.Println("right before", left, right, appereances)
			appereancesSum += left * appereances
			appereancesMap[left] = left * appereances
			appereances = 0

			if leftList.Len() == 0 {
				break
			}

			if left < right {
				oldLeft := left
				left = heap.Pop(leftList).(int)
				isSameNumber = oldLeft == left
				fmt.Println("removing left", left, right)
			} else {
				right = heap.Pop(rightList).(int)
				fmt.Println("removing right", left, right)
			}
		}
	}

	fmt.Println(appereancesMap)

	return appereancesSum
}
