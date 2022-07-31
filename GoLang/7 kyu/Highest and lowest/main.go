package main

import (
	"strings"
	"strconv"
	"fmt"
)

func main() {
	fmt.Println(HighAndLow("1 2 -3 4 5"))
}
  
func HighAndLow(in string) string {
	result := ""
	inSplit := strings.Fields(in)
	var res []int
	
	for i:= 0; i < len(inSplit); i++ {
	  intVar, _ := strconv.Atoi(string(inSplit[i]))
	  res = append(res, intVar)
	}
	
	max := res[0]
	min := res[0]
	for i := 0; i < len(res); i++ {
	  if max < res[i] { max = res[i] }
	  if res[i] < min { min = res[i] }
	}
	
	result = strconv.Itoa(max) + " " + strconv.Itoa(min)
	return result
}