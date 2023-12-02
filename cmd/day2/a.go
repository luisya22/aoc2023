package day2

import (
    "bufio"
    "github.com/luisya22/aoc2023/fileman"
    "github.com/spf13/cobra"
    "log"
)

var aCmd = &cobra.Command{
    Use: "a",
    Short: "Day 2, Problem A",
    Run: func(cmd *cobra.Command, args []string) {

        scanner, inputFile := fileman.GetFileLineBuffer("cmd/day2/input.txt")
        defer inputFile.Close()

        result := partA(scanner)

        log.Println("Result A: ", result)
    },
}

func partA(inputScanner *bufio.Scanner) int {
    return 0
}
