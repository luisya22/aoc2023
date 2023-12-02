package day2

import (
    "bufio"
    "github.com/luisya22/aoc2023/fileman"
    "github.com/spf13/cobra"
    "log"
)

var bCmd = &cobra.Command{
    Use: "b",
    Short: "Day 2, Problem B",
    Run: func(cmd *cobra.Command, args []string) {

        scanner, inputFile := fileman.GetFileLineBuffer("cmd/day2/input.txt")
        defer inputFile.Close()

        result := partB(scanner)

        log.Println("Result B: ", result)
    },
}

func partB(inputScanner *bufio.Scanner) int {
    return 0
}
