package day8

import (
	"bufio"
	"fmt"
	"log"
	"strings"

	"github.com/luisya22/aoc2023/fileman"
	"github.com/spf13/cobra"
)

var bCmd = &cobra.Command{
	Use:   "b",
	Short: "Day 8, Problem B",
	Run: func(cmd *cobra.Command, args []string) {

		scanner, inputFile := fileman.GetFileLineBuffer("cmd/day8/input.txt")
		defer inputFile.Close()

		result := partB(scanner)

		log.Println("Result B: ", result)
	},
}

func partB(inputScanner *bufio.Scanner) uint64 {

	inputScanner.Scan()
	instructions := inputScanner.Text()

	nodesNetwork, paths, err := parseNodesB(inputScanner)
	if err != nil {
		fmt.Println(err.Error())
		return 0
	}

	stepsTaken := 0
	instructionsPointer := 0
	found := 0
	for found < len(paths)*4 {

		if instructionsPointer == len(instructions) {
			instructionsPointer = 0
		}

		nextInstruction := string(instructions[instructionsPointer])

		for i := 0; i < len(paths); i++ {

			actualStep := paths[i].actualStep
			actualNode := nodesNetwork[actualStep]

			if nextInstruction == "L" {
				paths[i].actualStep = actualNode.left
			} else {
				paths[i].actualStep = actualNode.right
			}

			if string(paths[i].actualStep[2]) == "Z" {
				paths[i].stepsToZ = paths[i].steps
				paths[i].steps = 0
				found++
			}

			paths[i].steps++

		}

		stepsTaken++
		instructionsPointer++
	}

	return lcmPaths(paths)
}

type path struct {
	actualStep string
	steps      uint64
	stepsToZ   uint64
}

func parseNodesB(inputScanner *bufio.Scanner) (map[string]node, []path, error) {

	nodeNetwork := make(map[string]node)
	paths := []path{}

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

		if string(step[2]) == "A" {
			paths = append(paths, path{
				actualStep: step,
				steps:      0,
				stepsToZ:   0,
			})
		}

		nodeNetwork[n.step] = n
	}

	return nodeNetwork, paths, nil
}

func remove[T any](slice []T, s int) []T {
	return append(slice[:s], slice[s+1:]...)
}

func gcd(a, b uint64) uint64 {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}

	return a
}

func lcm(a, b uint64) uint64 {
	return a * b / gcd(a, b)
}

func lcmPaths(paths []path) uint64 {
	result := paths[0].stepsToZ
	for _, p := range paths[1:] {
		result = lcm(result, p.stepsToZ)
	}

	return result
}
