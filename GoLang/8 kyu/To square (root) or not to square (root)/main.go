package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(SquareOrSquareRoot([]int{10, 9, 37, 36, 1}))
}

func SquareOrSquareRoot(arr []int) []int{
  var res []int
  for _, value := range arr {
    
    sqrt := math.Sqrt(float64(value))
    
    if sqrt == math.Floor(sqrt) {
      res = append(res, int(sqrt))
    } else {
      res = append(res, value*value)
    }
    
  }
  
  return res
}