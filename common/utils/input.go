package utils

import (
	"bufio"
	"os"
	"strconv"
)

// LoadFile takes a filename and returns the two columns of integers as arrays.
func LoadFile(filename string) ([]int, []int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, nil, err
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)
	count := 0
	leftList := make([]int, 0)
	rightList := make([]int, 0)
	for scanner.Scan() {
		value, err := strconv.Atoi(scanner.Text())

		if err != nil {
			return nil, nil, err
		}

		if count%2 == 0 {
			leftList = append(leftList, value)
		} else {
			rightList = append(rightList, value)
		}
		count++
	}

	return leftList, rightList, nil
}
