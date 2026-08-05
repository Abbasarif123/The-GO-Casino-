package main

import (
	"encoding/json"
	"net/http"
)

//structs for webrequests and outgoing responses

// spinrequest defines the structure of the data we expect from FROM the browser
// json: -> these tags tell Go how to map JSON keys to specific go variables
type SpinRequest struct {
	Bet     uint `json:"bet"`
	Balance uint `json:"balance"`
}

// spin response defines the structure of the data we send BACK to the browser
type SpinResponse struct {
	Spin       [][]string `json:"spin"`       // 3x3 grid
	WinAmount  uint       `json:"winAmount"`  // how much the user won this spin
	NewBalance uint       `json:"newBalance"` // the user's updated balance
	Message    string     `json:"message"`    // status message for the UI
}

// handspin is the API endpoint function, it will run everytime the user clicks spin\
func handleSpin(w http.ResponseWriter, r *http.Request) {
	//safety check -> only allow POST requests since we are sending data to the server
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	//read JSON data sent from the browser and store it in the SpinRequest struct
	var req SpinRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil { //error handling
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	//validation to make sure bet < balance
	if req.Bet > req.Balance {
		//send JSON message to browser
		json.NewEncoder(w).Encode(SpinResponse{Message: "Bet cannot be larger than balance!!"})
		return
	}

	//GAME LOGIC
	newBalance := req.Balance - req.Bet

	//generate the 3x3 spin using GetSpin (in spin.go)
	spin := GetSpin(symbolArr, 3, 3)
	//check for win
	winninglines := checkWin(spin, multipliers)
	//calculate total winnings
	var totalWin uint = 0
	for _, multi := range winninglines {
		totalWin += multi * req.Bet
	}

	//add winnings
	newBalance += totalWin

	//PREPARE RESPONSE
	msg := "Better luck next time!"
	if totalWin > 0 {
		msg = "You won!"
	}

	//package all the results into the SpinResp
	resp := SpinResponse{
		Spin:       spin,
		WinAmount:  totalWin,
		NewBalance: newBalance,
		Message:    msg,
	}

	//inform browser that JSON is comin back and then send it
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
