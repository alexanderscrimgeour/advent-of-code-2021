package main

import (
	"adventofcode2021/day1"
	"adventofcode2021/day2"
	"flag"
	"fmt"
)

func main() {
	day := flag.Int("d", 1, "The day to run")
	flag.Parse()
	reader := NewReader(*day)
	input, err := reader.GetInputLines(*day)
	if err != nil {
		panic("Error reading input")
	}
	switch *day {
	case 1:
		answer, err := day1.SolvePartOne(input)
		if err != nil {
			panic(fmt.Sprintf("Error calculating day 1 part 1 solution: ", err))
		}
		fmt.Printf("\nAnswer: %d", answer)

		answer, err = day1.SolvePartTwo(input)
		if err != nil {
			panic(fmt.Sprintf("Error calculating day 1 part 2 solution: ", err))
		}
		fmt.Printf("\nAnswer: %d", answer)

	case 2:
		answer, err := day2.SolvePartOne(input)
		if err != nil {
			panic(fmt.Sprintf("Error calculating day 2 part 1 solution: ", err))
		}
		fmt.Printf("\nAnswer: %d", answer)

		answer, err = day2.SolvePartTwo(input)
		if err != nil {
			panic(fmt.Sprintf("Error calculating day 2 part 2 solution: ", err))
		}
		fmt.Printf("\nAnswer: %d", answer)
	case 3:
		fallthrough
	case 4:
		fallthrough
	case 5:
		fallthrough
	case 6:
		fallthrough
	case 7:
		fallthrough
	case 8:
		fallthrough
	case 9:
		fallthrough
	case 10:
		fallthrough
	case 11:
		fallthrough
	case 12:
		fallthrough
	case 13:
		fallthrough
	case 14:
		fallthrough
	case 15:
		fallthrough
	case 16:
		fallthrough
	case 17:
		fallthrough
	case 18:
		fallthrough
	case 19:
		fallthrough
	case 20:
		fallthrough
	case 21:
		fallthrough
	case 22:
		fallthrough
	case 23:
		fallthrough
	case 24:
		fallthrough
	case 25:
		fallthrough
	default:
		fmt.Println("No solution for day: ", *day)
	}

}
