package main

import (
	"log"
	"net/http"
	"os"
)

// global vars for the game state.
// map for our symbols
var symbols = map[string]uint{
	"A": 4,
	"B": 7,
	"C": 12,
	"D": 20,
} //the letters are the symbols and the numbers are their frequency

var multipliers = map[string]uint{
	"A": 20,
	"B": 10,
	"C": 5,
	"D": 2,
} //how many times of your bet you get on the basis of the symbols

var symbolArr []string

// check for winnings
func checkWin(spin [][]string, multipliers map[string]uint) []uint {
	lines := []uint{}

	for _, row := range spin {
		win := true
		checkSymbol := row[0]
		for _, symbol := range row[1:] {
			if checkSymbol != symbol {
				win = false
				break
			}
		}
		if win {
			lines = append(lines, multipliers[checkSymbol])
		} else {
			lines = append(lines, 0)
		}
	}
	return lines
}

func main() {
	//initialise symbol array once the server starts
	symbolArr = GenerateSymbolArray(symbols)

	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/", fs)
	http.HandleFunc("/api/spin", handleSpin)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server starting on port %s...\n", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
