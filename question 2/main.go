package main

import (
	"fmt"
	"strings"
)

func main() {
	var input string = "{ [ ( ] ) }"
	var res string = "Yes"

	inputTrimmed := strings.Replace(input, " ", "", -1)
	tempSlices := make([]rune, 0)

	for _, char := range inputTrimmed {
		if char == '(' || char == '[' || char == '{' {
			tempSlices = append(tempSlices, char)
		} else if char == ')' {
			if len(tempSlices) == 0 || tempSlices[len(tempSlices)-1] != '(' {
				res = "No"
				break
			}
			tempSlices = tempSlices[:len(tempSlices)-1]
		} else if char == ']' {
			if len(tempSlices) == 0 || tempSlices[len(tempSlices)-1] != '[' {
				res = "No"
				break
			}
			tempSlices = tempSlices[:len(tempSlices)-1]
		} else if char == '}' {
			if len(tempSlices) == 0 || tempSlices[len(tempSlices)-1] != '{' {
				res = "No"
				break
			}
			tempSlices = tempSlices[:len(tempSlices)-1]
		}
	}

	fmt.Println("Input : " + input)
	fmt.Println("Result : " + res)
}
