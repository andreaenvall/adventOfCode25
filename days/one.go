package one

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func DayOne() string {
	part1, part2 := FindCode(PuzzleInput2())
	return fmt.Sprintf("Part 1: %d, Part 2: %d", part1, part2)
}

func FindCode(input []string) (int, int) {
	position := 50
	zerosCountpart1 := 0
	zerosCountpart2 := 0
	for _, step := range input {
		direction := string(step[0])
		stepsArr := strings.Split(step, direction)

		stepsInt, _ := strconv.Atoi(stepsArr[1])
		switch direction {
		case "L":
			for range stepsInt {
				position -= 1

				// Reset position if less than 0
				if position < 0 {
					position = 99
				}
				// Check for part 2
				if position == 0 {
					zerosCountpart2++
				}
			}
		case "R":
			for range stepsInt {
				position += 1

				// Reset position if greater than 99
				if position > 99 {
					position = 0
				}
				// Check for part 2
				if position == 0 {
					zerosCountpart2++
				}
			}
		}
		if position == 0 {
			zerosCountpart1++
		}

	}

	return zerosCountpart1, zerosCountpart2
}

func PuzzleInput2() []string {
	var linearray []string
	f, _ := os.Open("days/puzzleinputday1.txt")

	r := bufio.NewReader(f)
	for {
		line, err := r.ReadString('\n')
		if err != nil {
			break
		}
		linearray = append(linearray, strings.TrimSpace(line))
	}
	return linearray
}
