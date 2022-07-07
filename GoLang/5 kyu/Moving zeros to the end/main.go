package main

import "fmt"

func MoveZeros(arr []int) []int {
	var nums []int
	var zeros []int
	
	for _, elem := range arr {
	  
	  if elem == 0 {
		zeros = append(zeros, 0)
	  } else {
		nums = append(nums, elem)
	  }
	  
	}
	
	return append(nums, zeros...)
}

func main() {
	fmt.Println(MoveZeros([]int{1, 2, 0, 5, 0, 8, 0, 3, 0, 1}))
}