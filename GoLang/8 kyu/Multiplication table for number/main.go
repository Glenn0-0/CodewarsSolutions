package main

import "fmt"

func main() {
	fmt.Println(MultiTable(4))
}

func MultiTable(number int) string {
	res := ""
	for i := 1; i <= 10; i++ {
	  if i != 10 {
		res += fmt.Sprintf("%v * %v = %v\n", i, number, i*number)
	  } else {
		res += fmt.Sprintf("%v * %v = %v", i, number, i*number)
	  }
	  
	} 
	return res
  }