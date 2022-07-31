package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(StringToNumber("-5"))
}

func StringToNumber(str string) int {
	res, _ := strconv.Atoi(str)    
	return res
  }