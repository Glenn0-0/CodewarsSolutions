package main

import (
	"fmt"
	"strings"
	"strconv"
)

func main() {
	fmt.Println(Points([]string{"1:0","2:0","3:0","4:0","2:1","3:1","4:1","3:2","4:2","4:3"}))
}
  
func Points(games []string) int {
	points := 0
	
	for _, score := range(games) {
	  nums := strings.Split(score, ":")
	  
	  team1, _ := strconv.Atoi(nums[0])
	  team2, _ := strconv.Atoi(nums[1])
	  
	  switch {
		case team1 > team2: points += 3
		case team1 == team2: points += 1
	  }
  
	}
	return points;
}