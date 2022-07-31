package main

import (
	"strings"
	"fmt"
)

func main() {
	fmt.Println(Solve("dgljirsljsefj"))
}

func delDuplicates(str string) string {
  var result string = ""
  for _, elem := range str {
    if strings.Count(result, string(elem)) == 0 {
      result += string(elem)
    }
  }
  return result
}

func occurIndexes(elem rune, str string) []int {
  var occur []int
  for index, value := range str {
    if elem == value {
      occur = append(occur, index)
    }
  }
  return occur
}

func min(a, b rune) rune {
  if a > b {return b}
  return a
}

func Solve(s string) rune {
  var letters string = delDuplicates(s)

  var maxDiff int = -1 
  var char rune  

  for _, value := range letters {
    var diff int 
    var occur []int = occurIndexes(value, s)

    var lastIndex int = occur[len(occur)-1]
    var firstIndex int = occur[0]
    diff = lastIndex - firstIndex

    if diff > maxDiff {
      maxDiff = diff
      char = value

    } else if diff == maxDiff { 
      char = min(char, value)
    }
  }
  
  return char
}