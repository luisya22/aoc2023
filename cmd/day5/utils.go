package day5

import (
	"fmt"
	"strconv"
	"strings"
)

type seedMapLine struct {
	source      uint64
	destination uint64
	mapRange    uint64
}

func addLine(sm []seedMapLine, line string) ([]seedMapLine, error) {
	seedMapLineStr := strings.Split(strings.TrimSpace(line), " ")

	destinationStart, err := strconv.ParseUint(seedMapLineStr[0], 10, 64)
	if err != nil {
		return sm, fmt.Errorf("error parsing destination: %v; %v", line, err)
	}

	sourceStart, err := strconv.ParseUint(seedMapLineStr[1], 10, 64)
	if err != nil {
		return sm, fmt.Errorf("error parsing source: %v; %v", line, err)
	}

	mapRange, err := strconv.ParseUint(seedMapLineStr[2], 10, 64)
	if err != nil {
		return sm, fmt.Errorf("error parsing mapRange: %v; %v", line, err)
	}

	seedML := seedMapLine{
		source:      sourceStart,
		destination: destinationStart,
		mapRange:    mapRange,
	}

	sm = append(sm, seedML)

	return sm, err
}

func getSeeds(seedsStr string) ([]uint64, error) {

	seedsArr := []uint64{}

	seedsSplt := strings.Split(seedsStr, ":")

	seeds := strings.Split(strings.TrimSpace(seedsSplt[1]), " ")

	for _, seedStr := range seeds {
		seed, err := strconv.ParseUint(seedStr, 10, 64)
		if err != nil {
			return []uint64{}, err
		}

		seedsArr = append(seedsArr, seed)
	}

	return seedsArr, nil
}

func transferSeeds(seedMap []seedMapLine, seeds []uint64) ([]uint64, error) {
	for seedKey, seed := range seeds {
		for _, seedMapLine := range seedMap {
			if seed >= seedMapLine.source && seed <= seedMapLine.source+seedMapLine.mapRange {
				diff := seed - seedMapLine.source

				newSeedNum := seedMapLine.destination + uint64(diff)

				seeds[seedKey] = newSeedNum

				break
			}
		}
	}

	return seeds, nil
}

func isDigit(d uint8) bool {
	if d >= '0' && d <= '9' {
		return true
	}

	return false
}
