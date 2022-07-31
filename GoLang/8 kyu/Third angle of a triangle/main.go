package main

import "fmt"

func main() {
	fmt.Println(OtherAngle(75, 80))
}

func OtherAngle(a int, b int) int {
    return 180 - (a + b)
}