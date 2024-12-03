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
	//fmt.Println(sumList)
	var sum int
	for _, value := range sumList {
		sum += value
	}
	fmt.Println(sum)

}
