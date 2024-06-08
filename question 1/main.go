package main

import (
	"fmt"
	"strings"
)

func main() {

	// Input
	// You can change the input here, to test if the program give a desired output
	var queries []int = []int{1, 8, 9, 3}
	var input string = "abbcccd"

	var result []string = []string{}

	// Create a temporary array to store all possible substring
	/**********Start************/
	var subsOfInput []string
	current := string(input[0])
	for i := 1; i < len(input); i++ {
		if input[i] == input[i-1] {
			subsOfInput = append(subsOfInput, current)
			current += string(input[i])
		} else {
			subsOfInput = append(subsOfInput, current)
			current = string(input[i])
		}
	}

	subsOfInput = append(subsOfInput, current)

	/**********END************/

	//Calculate weight for every possible substring
	/**********Start************/
	var weightedString []int
	for i := 0; i < len(subsOfInput); i++ {
		if len(subsOfInput[i]) > 1 {
			tempString := strings.ToLower(subsOfInput[i])
			tempTotal := 0
			for j := 0; j < len(tempString); j++ {
				tempTotal += int(tempString[j]) - 96
				if j == len(tempString)-1 {
					weightedString = append(weightedString, tempTotal)
					break
				}

			}
		} else if len(subsOfInput[i]) == 1 {
			weightedString = append(weightedString, int([]rune(strings.ToLower(subsOfInput[i]))[0])-96)
		}
	}
	/**********End************/

	//Loop around queries, assign  initial value of "No" to the result slice, and then
	//redeclare its value if there is a weight of substring that match the value of queries
	/**********Start************/
	for i := 0; i < len(queries); i++ {
		result = append(result, "No")
		for j := 0; j < len(weightedString); j++ {
			if queries[i] == weightedString[j] {
				result[i] = "Yes"
				break
			}
		}
	}
	/**********End************/

	fmt.Println("Input : " + input)
	fmt.Println("Queries : ")
	fmt.Println(queries)
	fmt.Println("Result : ")
	fmt.Println(result)

}
