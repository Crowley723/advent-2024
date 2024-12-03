package main

import "github.com/Crowley723/advent-2024/common/utils"

func main() {
	Day2Inputs, err := utils.LoadFile("input.txt", Parse)

	if err != nil {
		panic(err)
	}
}
