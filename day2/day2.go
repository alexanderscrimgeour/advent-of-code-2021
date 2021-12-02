package day2

import (
	"fmt"
	"strconv"
	"strings"
)

const (
	fwd  = "forward"
	up   = "up"
	down = "down"
)

type Coord struct {
	X, Y int
}

func moveHorizontal(coord Coord, val int) Coord {
	coord.X = coord.X + val
	return coord
}

func moveVertical(coord Coord, val int) Coord {
	coord.Y = coord.Y + val
	return coord
}

func SolvePartOne(input []string) (int, error) {
	loc := Coord{0, 0}

	for _, instruction := range input {
		args := strings.Split(instruction, " ")
		if len(args) > 2 {
			panic("instruction arguments contains too many parts")
		}
		cmd := args[0]
		val, err := strconv.Atoi(args[1])
		if err != nil {
			panic(err)
		}
		switch cmd {
		case fwd:
			loc = moveHorizontal(loc, val)
		case up:
			if loc.Y > -val {
				fmt.Printf("You're attempting to fly a submarine.")
			} else {
				loc = moveVertical(loc, val)
			}
		case down:
			loc = moveVertical(loc, -val)
		}
	}
	ans := loc.X * loc.Y
	return ans, nil
}

// func SolvePartTwo(input []string) (int, error) {

// 	return nil, nil
// }
