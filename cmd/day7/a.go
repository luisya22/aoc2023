package day7

import (
	"bufio"
	"fmt"
	"log"
	"sort"
	"strconv"
	"strings"

	"github.com/luisya22/aoc2023/fileman"
	"github.com/spf13/cobra"
)

var aCmd = &cobra.Command{
	Use:   "a",
	Short: "Day 7, Problem A",
	Run: func(cmd *cobra.Command, args []string) {

		scanner, inputFile := fileman.GetFileLineBuffer("cmd/day7/input.txt")
		defer inputFile.Close()

		result := partA(scanner)

		log.Println("Result A: ", result)
	},
}

func partA(inputScanner *bufio.Scanner) int {

	playsSum := 0
	plays := []play{}

	for inputScanner.Scan() {
		line := inputScanner.Text()

		lineSplit := strings.Split(strings.TrimSpace(line), " ")

		hand := lineSplit[0]
		bid, err := strconv.Atoi(lineSplit[1])
		if err != nil {
			fmt.Println(err.Error())
			return -1
		}

		hType, err := getHandType(hand)
		if err != nil {
			fmt.Println(err.Error())
			return -1
		}

		p := play{
			hand:  hand,
			bid:   bid,
			hType: hType,
		}

		plays = append(plays, p)
	}

	sort.Sort(byHandType(plays))

	for i := 0; i < len(plays); i++ {
		p := plays[i]
		fmt.Printf("Points: %v - %v - %v -> %v - %v = %v\n", gt(p.hType), i+1, p.bid, p.hand, (i+1)*p.bid, playsSum)
		playsSum += (i + 1) * p.bid
	}

	return playsSum
}

type byHandType []play

func (ht byHandType) Len() int {
	return len(ht)
}

func (ht byHandType) Swap(i, j int) {
	ht[i], ht[j] = ht[j], ht[i]
}

func (ht byHandType) Less(i, j int) bool {

	if ht[i].hType != ht[j].hType {
		return ht[i].hType < ht[j].hType
	}

	// For each validate one to one
	for x := 0; x < len(ht[i].hand); x++ {
		iValue := getCardValue(string(ht[i].hand[x]))
		jValue := getCardValue(string(ht[j].hand[x]))

		if iValue == jValue {
			continue
		} else {
			return iValue < jValue
		}
	}

	return false
}

func getCardValue(card string) int {
	cardMap := map[string]int{
		"A": 14,
		"K": 13,
		"Q": 12,
		"J": 11,
		"T": 10,
		"9": 9,
		"8": 8,
		"7": 7,
		"6": 6,
		"5": 5,
		"4": 4,
		"3": 3,
		"2": 2,
		"1": 1,
	}

	return cardMap[card]
}

func getHandType(hand string) (handType, error) {

	if len(hand) != 5 {
		return 0, fmt.Errorf("incorrect hand length")
	}

	handMap := make(map[string]int)

	for i := 0; i < len(hand); i++ {
		card := string(hand[i])

		handMap[card]++
	}

	containsMap := make(map[int]struct{})
	for _, val := range handMap {
		_, containsPair := containsMap[2]
		_, containsThree := containsMap[3]

		if val == 5 {
			return fiveOfKind, nil
		}

		if val == 4 {
			return fourOfKind, nil
		}

		if val == 3 && containsPair {
			return fullHouse, nil
		}

		if val == 2 && containsThree {
			return fullHouse, nil
		}

		if val == 2 && containsPair {
			return twoPair, nil
		}

		containsMap[val] = struct{}{}
	}

	_, containsPair := containsMap[2]
	if containsPair {
		return onePair, nil
	}

	_, containsThree := containsMap[3]
	if containsThree {
		return threeOfKind, nil
	}

	return highCard, nil

}
