package days

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func DayTwo() string {
	fmt.Println(PuzzleInput2())
	part1 := FindCodeDay2(PuzzleInput2())
	return fmt.Sprintf("Part 1: %s", part1)
}

func FindCodeDay2(input string) string {
	// Split the input into lines
	var invalidIDs []int
	productIDs := strings.Split(input, ",")
	for _, v := range productIDs {
		productIDs = strings.Split(v, "-")
		isInValidCound := IsInvalidInput(productIDs[0], productIDs[1])
		invalidIDs = append(invalidIDs, isInValidCound...)
	}
	var result int
	for _, id := range invalidIDs {
		result += id
	}

	return fmt.Sprintf("Invalid IDs: %d", result)
}

func IsInvalidInput(startID, endID string) []int {
	var invalidIDs []int
	startIDint, err := strconv.Atoi(startID)
	if err != nil {
		return invalidIDs
	}
	endIDint, err := strconv.Atoi(endID)
	if err != nil {
		return invalidIDs
	}

	for {
		lenOfSequence := digitCount(startIDint)

		if lenOfSequence > 1 {
			isSameigits := checkSameDigits(startIDint)
			if isSameigits {
				fmt.Println("appending invalid id isSameigits ", startIDint)
				invalidIDs = append(invalidIDs, startIDint)
				startIDint++
				continue
			}
		}

		if lenOfSequence%2 != 0 {
			if startIDint == endIDint {
				fmt.Println("Ending")
				break
			}
			invalidID := isInvalidID(lenOfSequence, startIDint)
			if invalidID {
				fmt.Println("appending invalid id", startIDint)
				invalidIDs = append(invalidIDs, startIDint)
			}

			startIDint++
			continue
		}
		half := lenOfSequence / 2

		if fmt.Sprint(startIDint)[:half] == fmt.Sprint(startIDint)[half:] {
			fmt.Println("appending invalid id", startIDint)
			invalidIDs = append(invalidIDs, startIDint)
		} else {
			invalidID := isInvalidIDtwoNum(startIDint)
			if invalidID {
				fmt.Println("appending invalid id from isInvalidIDtwoNum", startIDint)
				invalidIDs = append(invalidIDs, startIDint)
			}
		}

		if startIDint == endIDint {
			fmt.Println("Ending")
			break
		}
		startIDint++
	}

	return invalidIDs
}

func isInvalidID(lenOfSequence, startID int) bool {
	third := lenOfSequence / 3
	if fmt.Sprint(startID)[:third] == fmt.Sprint(startID)[third:third*2] && fmt.Sprint(startID)[:third] == fmt.Sprint(startID)[third*2:] {
		return true
	}
	return false
}

func isInvalidIDtwoNum(startID int) bool {
	idString := fmt.Sprint(startID)

	if len(idString) < 3 {
		return false
	}
	start := 0
	position := 2
	for {
		if fmt.Sprint(startID)[start:position] == fmt.Sprint(startID)[position:position+2] {
			position += 2
			start += 2
		} else {
			return false
		}

		if position == len(idString) {
			break
		}

	}

	return true
}

func digitCount(num int) int {
	var count int = 0
	for num > 0 {
		num = num / 10
		count = count + 1
	}
	return count
}

func checkSameDigits(n int) bool {
	digit := n % 10

	for n != 0 {
		currentDigit := n % 10
		n /= 10

		if currentDigit != digit {
			return false
		}
	}

	return true
}

func PuzzleInput2() string {
	var linearray string
	f, _ := os.Open("days/puzzleinputday2.txt")

	r := bufio.NewReader(f)
	for {
		line, err := r.ReadString('\n')
		if err != nil {
			break
		}
		linearray += strings.TrimSpace(line)
	}
	return linearray
}

// 73694270688
