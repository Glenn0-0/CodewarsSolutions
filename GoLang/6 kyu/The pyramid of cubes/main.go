package main

import "fmt"

func main() {
	fmt.Println(FindHeight(32))
}

func FindHeight(cubes int) int {
	count := 0
	layer := -1
	for cubes >= count {
	  cubes -= count
	  count += layer + 2
  
	  layer++
  
	}
	return layer
}