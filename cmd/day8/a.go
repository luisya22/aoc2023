package day8

import (
	"bufio"
	"fmt"
	"log"
	"strings"

	"github.com/luisya22/aoc2023/fileman"
	"github.com/spf13/cobra"
)

var aCmd = &cobra.Command{
	Use:   "a",
	Short: "Day 8, Problem A",
	Run: func(cmd *cobra.Command, args []string) {

		scanner, inputFile := fileman.GetFileLineBuffer("cmd/day8/input.txt")
		defer inputFile.Close()

		result := partA(scanner)

		log.Println("Result A: ", result)
	},
}

func partA(inputScanner *bufio.Scanner) int {

	inputScanner.Scan()
	instructions := inputScanner.Text()

	nodesNetwork, err := parseNodes(inputScanner)
	if err != nil {
		fmt.Println(err.Error())
		return -1
	}

	stepsTaken := 0
	actualStep := "AAA"
	instructionsPointer := 0
	for actualStep != "ZZZ" {

		if instructionsPointer == len(instructions) {
			instructionsPointer = 0
		}

		nextInstruction := string(instructions[instructionsPointer])
		actualNode := nodesNetwork[actualStep]

		if nextInstruction == "L" {
			actualStep = actualNode.left
		} else {
			actualStep = actualNode.right
		}

		stepsTaken++
		instructionsPointer++
	}

	return stepsTaken
}

func parseNodes(inputScanner *bufio.Scanner) (map[string]node, error) {

	nodeNetwork := make(map[string]node)

	counter := 0
	for inputScanner.Scan() {
		line := inputScanner.Text()

		if line == "" {
			continue
		}

		lineSplt := strings.Split(strings.ReplaceAll(line, " ", ""), "=")

		step := lineSplt[0]
		trimParenthesis := strings.ReplaceAll(strings.ReplaceAll(lineSplt[1], "(", ""), ")", "")
		nextSteps := strings.Split(trimParenthesis, ",")

		n := node{
			step:  step,
			left:  nextSteps[0],
			right: nextSteps[1],
		}

		counter++

		nodeNetwork[n.step] = n
	}

	return nodeNetwork, nil
}
