package main

import "fmt"

func main() {
	fmt.Println(multipleOfIndex([]int{22, -6, 32, 82, 9, 25}))
}

func multipleOfIndex (ints []int) (res []int) {
	for index := 1; index < len(ints); index++ {
	  
	  elem := ints[index]
	  
	  if elem%index == 0 {
		res = append(res, elem)
	  }
	  
	}
	return res
  }