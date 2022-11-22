package main

import (
	"fmt"
	"math/rand"
)

var beats map[string]string
var possibilities [6]string

func input(prompt string) string {
	var x string
	fmt.Print(prompt)
	fmt.Scanln(&x)
	return x
}

func main() {
	possibilities := [6]string{"rock", "paper", "scissors"}
	beats := make(map[string]string)
	beats["rock"] = "scissors"
	beats["scissors"] = "paper"
	beats["paper"] = "rock"

	for {
		choice := input("rock paper or scissors? ")
		comp_choice := possibilities[rand.Intn(3)]
		// weed out bad inputs
		if choice != "rock" && choice != "paper" && choice != "scissors" {
			continue
		}
		if beats[comp_choice] == choice {
			fmt.Println("You lose, the computer chose " + comp_choice)
			continue
		}
		if beats[choice] == comp_choice {
			fmt.Println("You win, the computer chose " + comp_choice)
			continue
		}
		if comp_choice == choice {
			fmt.Println("Tie, the computer also chose " + comp_choice)
			continue
		}
	}
}
