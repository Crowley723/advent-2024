package main

import (
	"bufio"
	"os"
)

func Parse(file *os.File) (Day3Inputs, error) {
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	rows := make([]string, 0)
	for scanner.Scan() {
		rows = append(rows, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return Day3Inputs{}, err
	}

	return Day3Inputs{Rows: rows}, nil
}
