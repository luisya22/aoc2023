package day9

import (
	"strconv"
	"strings"
)

type historyPredictions [][]int

func parseData(line string) (historyPredictions, error) {

	historyData := []int{}
	numsStr := strings.Split(strings.TrimSpace(line), " ")

	for _, num := range numsStr {
		num, err := strconv.Atoi(num)
		if err != nil {
			return [][]int{}, err
		}

		historyData = append(historyData, num)
	}

	return [][]int{historyData}, nil
}
