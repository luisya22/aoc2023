package day3

import "bufio"

type engineSchematic struct {
	lines   []string
	symbols map[position]*symbol
}

type position struct {
	y int
	x int
}

type symbol struct {
	position position
	kind     string
	parts    []int
}

func isAdjacent(pos position, symbols map[position]*symbol) (bool, []position) {
	adjacentPosToCheck := [][2]int{
		{-1, -1},
		{-1, 0},
		{-1, 1},
		{0, -1},
		{0, 1},
		{1, -1},
		{1, 0},
		{1, 1},
	}

	adjacentPos := []position{}

	for _, ap := range adjacentPosToCheck {
		checkPos := position{
			y: pos.y + ap[0],
			x: pos.x + ap[1],
		}

		_, foundSymbol := symbols[checkPos]
		if foundSymbol {
			adjacentPos = append(adjacentPos, checkPos)
		}
	}

	return len(adjacentPos) > 0, adjacentPos
}

func (es *engineSchematic) addPart(partNum int, pos position) {

	s, ok := es.symbols[pos]
	if ok {
		s.parts = append(s.parts, partNum)
	}
}

func parseEngine(scanner *bufio.Scanner) (engineSchematic, error) {

	es := engineSchematic{
		lines:   []string{},
		symbols: make(map[position]*symbol),
	}

	lineNumber := 0
	for scanner.Scan() {
		line := scanner.Text()

		es.lines = append(es.lines, line)

		for i := 0; i < len(line); i++ {
			if !isDigit(line[i]) && string(line[i]) != "." {
				pos := position{
					y: lineNumber,
					x: i,
				}
				s := &symbol{
					position: pos,
					kind:     string(line[i]),
					parts:    []int{},
				}

				es.symbols[pos] = s
			}
		}

		lineNumber++

	}

	return es, nil
}

func isDigit(d uint8) bool {
	if d >= '0' && d <= '9' {
		return true
	}

	return false
}
