package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(Solve("1234", 2))
}

func Solve(st string, k int) int {
  lenOfNum := len(st) - k
  maxNum := 0
  
  for i := 0; i <= len(st) - lenOfNum; i++ {
    
    currentNumStr := st[i : i + lenOfNum]
    numInt, _ := strconv.Atoi(currentNumStr)
    
    if numInt > maxNum { 
      maxNum = numInt 
    }
    
  }
  
  return maxNum
}