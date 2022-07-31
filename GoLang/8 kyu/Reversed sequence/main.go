package main

import "fmt"

func main() {
	fmt.Println(ReverseSeq(5))
}

func ReverseSeq(n int) []int {
	var res []int
	for i := n; i > 0; i-- {
	  res = append(res, i)
	}
	return res
  }