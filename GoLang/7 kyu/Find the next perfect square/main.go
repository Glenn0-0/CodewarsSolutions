package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(FindNextSquare(49))
}

func FindNextSquare(sq int64) int64 {
  num := math.Sqrt(float64(sq))
  if num == math.Floor(num) {
    return int64((num + 1)*(num + 1))
  }
  return -1
}