package main

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/Crowley723/advent-2024/common/utils"
)

func main() {
	Day3Inputs, err := utils.LoadFile("input.txt", Parse)
	if err != nil {
		panic(err)
	}

	validOps := regexp.MustCompile(`(mul\(\d+,\d+\)|do\(\)|don't\(\))`)
	mulRe := regexp.MustCompile(`mul\((\d+),(\d+)\)`)
	doRe := regexp.MustCompile(`do\(\)`)
	dontRe := regexp.MustCompile(`don't\(\)`)

	multOperations := make([]Operation, 0)

	for _, row := range Day3Inputs.Rows {
		validMatches := validOps.FindAllString(row, -1)
		for _, match := range validMatches {
			if value := mulRe.FindAllStringSubmatch(match, -1); value != nil {
				for _, op := range value {
					num1, _ := strconv.Atoi(op[1])
					num2, _ := strconv.Atoi(op[2])
					multOperations = append(multOperations, Operation{MUL, num1, num2})
					fmt.Printf("%d * %d\n", num1, num2)
				}
			} else if doRe.MatchString(match) {
				multOperations = append(multOperations, Operation{DO, 0, 0})
				fmt.Printf("do\n")
			} else if dontRe.MatchString(match) {
				multOperations = append(multOperations, Operation{DONT, 0, 0})
				fmt.Printf("dont\n")
			}

		}
	}
	var (
		disableOperations bool = false
		totalSum          int  = 0
	)
	for _, op := range multOperations {
		if op.Operation == DONT {
			disableOperations = true
		} else if op.Operation == DO {
			disableOperations = false
		}
		if !disableOperations && op.Operation == MUL {
			totalSum += op.op1 * op.op2
		}
	}
	fmt.Printf("Total Sum: %d\n", totalSum)

}
