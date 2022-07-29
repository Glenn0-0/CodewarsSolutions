package main

import "fmt"

func main() {
	fmt.Println(ReverseList([]int{1,4,7,2,4,9,0,3}))
}

func ReverseList(lst []int) []int {
	res := []int{}
	for i := len(lst) - 1; i >= 0; i-- {
	  res = append(res, lst[i])
	}
	return res
  }