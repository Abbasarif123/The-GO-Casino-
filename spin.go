package main

import (
	"math/rand"
)

// construct an array of symbols occuring as per their frequencies and select a random index from that array
// generates slice on the basis of the symbols map
func GenerateSymbolArray(symbols map[string]uint) []string {
	symbolArr := []string{}
	//symbol -> key, count ->value
	for symbol, count := range symbols {
		for i := uint(0); i < count; i++ {
			symbolArr = append(symbolArr, symbol)
		}
	}
	return symbolArr
}

// random number generator
func GetRandomNumber(min int, max int) int {
	randomNumber := rand.Intn(max-min+1) + min
	return randomNumber
}

// create a spin
// generate the columns and construct the 2d slice
func GetSpin(reel []string, rows int, cols int) [][]string {
	result := [][]string{}

	for i := 0; i < rows; i++ {
		result = append(result, []string{}) //insert empty rows
	}

	for col := 0; col < cols; col++ {
		selected := map[int]bool{} //map tracks which positions have already been selected
		for row := 0; row < rows; row++ {
			for true {
				randomIndex := GetRandomNumber(0, len(reel)-1)
				_, exists := selected[randomIndex]
				if !exists { //if it doesnt exist then mark that position as true and put that index in the result
					selected[randomIndex] = true
					result[row] = append(result[row], reel[randomIndex])
					break
				}
			}
		}
	}
	return result
}

// spin printer
// func PrintSpin(spin [][]string) {
// 	for _, row := range spin {
// 		for j, symbol := range row {
// 			fmt.Printf(symbol)
// 			if j != len(row)-1 {
// 				fmt.Printf(" | ")
// 			}
// 		}
// 		fmt.Println()
// 	}
// }
//NO LONGER NEEDED FOR WEB TOOL
