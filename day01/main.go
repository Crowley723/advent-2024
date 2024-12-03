package main

import (
	"fmt"
	"sort"

	"github.com/Crowley723/advent-2024/common/utils"
)

func main() {
	leftList, rightList, err := utils.LoadFile("input.txt")

	if err != nil {
		panic(err)
	}

	sort.Ints(leftList)
	sort.Ints(rightList)

	sumList := make([]int, 0)

	for index, leftValue := range leftList {
		difference := utils.Abs(leftValue - rightList[index])
		sumList = append(sumList, difference)
	}
	sort.Ints(sumList)
	var sum int
	for _, value := range sumList {
		sum += value
	}
	fmt.Printf("The sum of the differences (Part 1) is: %d\n", sum)
	//END PART 1 -- START PART 2

	rightOccurrenceMap := utils.MapOccurrences(rightList)
	sum = 0

	for _, value := range leftList {
		sum += value * rightOccurrenceMap[value]
	}
	fmt.Printf("The similarity score (Part 2) is: %d\n", sum)

}
