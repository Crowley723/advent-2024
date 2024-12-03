package main

import (
	"fmt"
	"sort"

	"github.com/Crowley723/advent-2024/common/utils"
)

func main() {
	Day1Inputs, err := utils.LoadFile("input.txt", Parse)

	if err != nil {
		panic(err)
	}

	sort.Ints(Day1Inputs.LeftList)
	sort.Ints(Day1Inputs.RightList)

	sumList := make([]int, 0)

	for index, leftValue := range Day1Inputs.LeftList {
		difference := utils.Abs(leftValue - Day1Inputs.RightList[index])
		sumList = append(sumList, difference)
	}
	sort.Ints(sumList)
	var sum int
	for _, value := range sumList {
		sum += value
	}
	fmt.Printf("The sum of the differences (Part 1) is: %d\n", sum)
	//END PART 1 -- START PART 2

	rightOccurrenceMap := utils.MapOccurrences(Day1Inputs.RightList)
	sum = 0

	for _, value := range Day1Inputs.LeftList {
		sum += value * rightOccurrenceMap[value]
	}
	fmt.Printf("The similarity score (Part 2) is: %d\n", sum)

}
