package main

import "fmt"

func main() {
	fmt.Println(QuarterOf(4))
}

func QuarterOf(month int) int {
	switch month {
	  case 1, 2, 3: return 1
	  case 4, 5, 6: return 2
	  case 7, 8, 9: return 3
	  case 10, 11, 12: return 4
	}
	return 0
}