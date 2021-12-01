package day1

import "strconv"

func SolvePartOne(input []string) (int, error) {
	count := 0

	var last int
	for i, val := range input {
		intVal, _ := strconv.Atoi(val)
		if i == 0 {
			last = intVal
			continue
		}
		if intVal > last {
			count++
		}
		last = intVal
	}
	return count, nil
}

func SolvePartTwo(input []string) (int, error) {
	count := 0
	windowSize := 3
	var last int
	for i := 0; i < len(input); i++ {
		intVal, _ := strconv.Atoi(input[i])
		// The first <windowSize> elements are simply summed
		if i < windowSize {
			last += intVal
			continue
		}
		sum := 0
		for j := 0; j < windowSize; j++ {
			val, _ := strconv.Atoi(input[i-j])
			sum += val
		}
		if sum > last {
			count++
		}
		last = sum
	}
	return count, nil
}
