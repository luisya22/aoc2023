package day2

import (
	"strconv"
	"strings"
)

type game struct {
	id        int
	gameCubes []gameCube
}

type gameCube struct {
	cubes []cube
}

type cube struct {
	color    string
	quantity int
}

func parseGame(input string) (*game, error) {

	splitStr := strings.Split(input, ":")

	// Get Game Id
	idStr := strings.TrimPrefix(splitStr[0], "Game ")

	gameId, err := strconv.Atoi(idStr)
	if err != nil {
		return nil, err
	}

	// Get gamecubes
	gCubesStr := strings.Split(splitStr[1], ";")
	gameCubes := []gameCube{}

	for _, gc := range gCubesStr {
		gameCubeData := gameCube{}

		cubes := strings.Split(gc, ",")
		for _, c := range cubes {
			dataStr := strings.Split(strings.TrimSpace(c), " ")

			quantity, err := strconv.Atoi(dataStr[0])
			if err != nil {
				return nil, err
			}

			color := dataStr[1]

			cubeData := cube{
				color:    color,
				quantity: quantity,
			}

			gameCubeData.cubes = append(gameCubeData.cubes, cubeData)

		}

		gameCubes = append(gameCubes, gameCubeData)
	}

	gameData := &game{
		id:        gameId,
		gameCubes: gameCubes,
	}

	return gameData, nil
}
