package main

import (
	"bufio"
	"os"
	"strconv"
)

type Day1Inputs struct {
	LeftList  []int
	RightList []int
}

func Parse(file *os.File) (Day1Inputs, error) {
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)
	count := 0
	leftList := make([]int, 0)
	rightList := make([]int, 0)
	for scanner.Scan() {
		value, err := strconv.Atoi(scanner.Text())

		if err != nil {
			return Day1Inputs{}, err
		}

		if count%2 == 0 {
			leftList = append(leftList, value)
		} else {
			rightList = append(rightList, value)
		}
		count++
	}
	if err := scanner.Err(); err != nil {
		return Day1Inputs{}, err
	}

	return Day1Inputs{leftList, rightList}, nil
}
