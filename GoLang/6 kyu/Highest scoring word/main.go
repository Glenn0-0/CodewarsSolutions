package main

import (
	"fmt"
	"strings"
)

func scoreCount(word string) int {
	var score = 0
	points := map[string]int {"a": 1,"b": 2,"c": 3,"d": 4,"e": 5,"f": 6,"g": 7,"h": 8,"i": 9,"j": 10,"k": 11,"l": 12,"m": 13,"n": 14,"o": 15,"p": 16,"q": 17,"r": 18,"s": 19,"t": 20,"u": 21,"v": 22,"w": 23,"x": 24,"y": 25,"z": 26,}
	for _, letter := range word {
	  score += points[string(letter)]
	}
	return score
}
  
func High(s string) string {
	var minIndex int = 0
	var maxScore int = 0
	
	words := strings.Split(s, " ")
	for index, word := range words {
	  score := scoreCount(word)

	  if score > maxScore {
		minIndex = index
		maxScore = score

	  } else if score == maxScore {

		if index < minIndex {
		  minIndex = index
		}
		
	  }
	}

	return words[minIndex]
}

func main() {
	fmt.Println(High("man i need a taxi up to ubud"))
}