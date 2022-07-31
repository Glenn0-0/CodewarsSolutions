package main

import "fmt"

func main() {
	fmt.Println(Rps("rock", "scissors"))
}

func Rps(p1, p2 string) string {
	result := "Player 2 won!"
	if (p1 == "scissors" && p2 == "paper") || (p1 == "rock" && p2 == "scissors") || (p1 == "paper" && p2 == "rock"){
	  result = "Player 1 won!"
	} else if p1 == p2 {
	  result = "Draw!"
	}
	return result
  }