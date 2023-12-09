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

var bCmd = &cobra.Command{
	Use:   "b",
	Short: "Day 7, Problem B",
	Run: func(cmd *cobra.Command, args []string) {

		scanner, inputFile := fileman.GetFileLineBuffer("cmd/day7/input.txt")
		defer inputFile.Close()

		result := partB(scanner)

		log.Println("Result B: ", result)
	},
}

func partB(inputScanner *bufio.Scanner) int {

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

		hType, err := getHandTypeB(hand)
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

	sort.Sort(byHandTypeB(plays))

	for i := 0; i < len(plays); i++ {
		p := plays[i]

		playsSum += (i + 1) * p.bid
	}

	return playsSum
}

type byHandTypeB []play

func (ht byHandTypeB) Len() int {
	return len(ht)
}

func (ht byHandTypeB) Swap(i, j int) {
	ht[i], ht[j] = ht[j], ht[i]
}

func (ht byHandTypeB) Less(i, j int) bool {

	if ht[i].hType != ht[j].hType {
		return ht[i].hType < ht[j].hType
	}

	// For each validate one to one
	for x := 0; x < len(ht[i].hand); x++ {
		iValue := getCardValueB(string(ht[i].hand[x]))
		jValue := getCardValueB(string(ht[j].hand[x]))

		if iValue == jValue {
			continue
		} else {
			return iValue < jValue
		}
	}

	return false
}

func getCardValueB(card string) int {
	cardMap := map[string]int{
		"A": 13,
		"K": 12,
		"Q": 11,
		"T": 10,
		"9": 9,
		"8": 8,
		"7": 7,
		"6": 6,
		"5": 5,
		"4": 4,
		"3": 3,
		"2": 2,
		"J": 1,
	}

	return cardMap[card]
}

func handleJ(handMap map[string]int) (handType, error) {
	containsJ := handMap["J"]

	// 5 J
	if containsJ == 5 {
		return fiveOfKind, nil
	}
	// 4 J
	if containsJ == 4 {
		return fiveOfKind, nil
	}
	// 3 J
	// JJJKK, JJJAK, JJJ2
	if containsJ == 3 {
		for key, val := range handMap {
			if key == "J" {
				continue
			}

			if val == 2 {
				return fiveOfKind, nil
			}

			if val == 1 {
				return fourOfKind, nil
			}
		}
	}
	// 2 J
	// JJAAA, JJAAK, JJAK2
	if containsJ == 2 {
		for key, val := range handMap {
			if key == "J" {
				continue
			}

			if val == 3 {
				return fiveOfKind, nil
			}

			if val == 2 {
				return fourOfKind, nil
			}
		}

		return threeOfKind, nil
	}
	// 1 J
	// JAAAA, JAAAK, JAA23, JA234, J2233
	if containsJ == 1 {
		containsMap := make(map[int]struct{})
		for key, val := range handMap {
			if key == "J" {
				continue
			}

			_, containsPair := containsMap[2]

			if val == 4 {
				return fiveOfKind, nil
			}

			if val == 3 {
				return fourOfKind, nil
			}

			if val == 2 && containsPair {
				return fullHouse, nil
			}

			containsMap[val] = struct{}{}
		}

		_, containsPair := containsMap[2]

		if containsPair {
			return threeOfKind, nil
		}

		return onePair, nil
	}

	return onePair, nil
}

func getHandTypeB(hand string) (handType, error) {

	if len(hand) != 5 {
		return 0, fmt.Errorf("incorrect hand length")
	}

	handMap := make(map[string]int)

	for i := 0; i < len(hand); i++ {
		card := string(hand[i])

		handMap[card]++
	}

	if strings.Contains(hand, "J") {
		return handleJ(handMap)
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
