package main

import "fmt"

func main() {
	fmt.Println(PowersOfTwo(3))
}

func powOfTwo(pow int) uint64 {
	var two uint64 = 1
	for i := 0; i < pow; i++ {
	  two *= 2
	}
	return two
  }
  
func PowersOfTwo(n int) (arr []uint64) {
	for i := 0; i <= n; i ++ {
	  arr = append(arr, powOfTwo(i))
	}
	return arr
}