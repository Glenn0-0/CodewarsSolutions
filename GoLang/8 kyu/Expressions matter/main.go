package main

import "fmt"

func main() {
	fmt.Println(ExpressionMatter(3, 4, 5))
}

func getMax(nums ...int) int {
	max := nums[0]
	for _, value := range nums {
	  if value > max { max = value }
	}
	return max
}
  
func ExpressionMatter(a int, b int, c int) int {
	v1 := a * b * c
	v2 := a + b * c
	v3 := a * b + c
	v4 := a + b + c
	v5 := (a + b) * c
	v6 := a * (b + c)
	max := getMax(v1, v2, v3, v4, v5, v6)
	return max
}