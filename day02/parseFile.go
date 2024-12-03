package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func Parse(file *os.File) (Day2Inputs, error) {
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	count := 0
	grid := make([][]int, 0)
	for scanner.Scan() {
		slice := strings.Fields(scanner.Text())

		for index, value := range slice {
			value, err := strconv.Atoi(string(value))
			if err != nil {
				return Day2Inputs{}, err
			}
			grid[count][index] = value
		}
	}
	if err := scanner.Err(); err != nil {
		return Day2Inputs{}, err
	}

	return Day2Inputs{Grid: grid}, nil
}
