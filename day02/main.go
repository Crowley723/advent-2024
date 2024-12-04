package main

import (
	"fmt"

	"github.com/Crowley723/advent-2024/common/utils"
)

func main() {
	Day2Inputs, err := utils.LoadFile("input.txt", Parse)

	if err != nil {
		panic(err)
	}
	var totalSafe int
	for _, row := range Day2Inputs.Grid {
		if validateRow(row) {
			totalSafe++
		}
	}
	fmt.Printf("Total Safe Reports: %d\n", totalSafe)
}

func validateRow(row []int) bool {
	var trend int = 0
	for colIndex := 0; colIndex < len(row)-1; colIndex++ {
		diff := utils.Abs(row[colIndex] - row[colIndex+1])
		if diff < 1 || diff > 3 {
			return false
		}
		if row[colIndex] > row[colIndex+1] {
			if trend == 1 {
				return false
			}
			trend = -1
		} else if row[colIndex] < row[colIndex+1] {
			if trend == -1 {
				return false
			}
			trend = 1
		}
	}
	return true
}
