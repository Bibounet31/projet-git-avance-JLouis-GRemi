package utils

import (
	"fmt"
	"strings"
)

// this function is used to make everything look smoother by separating
// different block of print
func printPolish() {
	fmt.Println("================")
}

// this function is used to verif at the start of every loop if the player is polite or not
// by answering "bonjour"
func bonjourCon(answer string) string {
	verifSalute := "bonjour"
	answer = strings.ToLower(answer)
	if answer == "" {
		return "That's not polite you have to say bonjour"
	} else if answer == verifSalute {
		return "Bonjour !"
	} else {
		return "Error"
	}
}
