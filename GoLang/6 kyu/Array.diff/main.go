package main

import "fmt"

func notInArray(array []int, elem int) bool {
	for _, value := range array {
		if value == elem {
			return false
		}
	}
	return true
}

func ArrayDiff(a, b []int) []int {
	var result []int
	for _, elem := range a {
		if notInArray(b, elem) {
			result = append(result, elem)
		}
	}
	return result
}

func main() {
	fmt.Println(ArrayDiff([]int{1, 2, 2, 4, 5, 5, 5, 1, 7, 0, 9, 0, 2, 3}, []int{2, 5, 9}))
}
