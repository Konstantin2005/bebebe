package main

import (
	"strings"
	"unicode"
)

func OneCalculate(str string) int {
	str = strings.ReplaceAll(str, " ", "")
	arr := []int{}

	curValue := 0
	preOp := '+'

	for i, c := range str {
		if i == len(str)-1 || unicode.IsDigit(c) {
			curValue = curValue*10 + (int(c) - '0')
		}

		if i == len(str)-1 || !unicode.IsDigit(c) {
			switch preOp {
			case '*':
				arr[len(arr)-1] *= curValue
			case '/':
				arr[len(arr)-1] /= curValue
			case '+':
				arr = append(arr, curValue)
			case '-':
				curValue = -curValue
				arr = append(arr, curValue)
			}
			curValue = 0
			preOp = c
		}
	}

	result := 0
	for _, item := range arr {
		result += item
	}

	return result
}
