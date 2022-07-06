package main

import "fmt"

func createRow(count, num int) []int {
	var row []int
	temp := count
	for i := 0; i < num; i++ {
	  row = append(row, count)
	  count += temp
	}
	return row
}

func MultiplicationTable(size int) [][]int {
	var table [][]int
	for i := 1; i <= size; i++ {
		table = append(table, createRow(i, size))
	  }
	return table
}

func main() {
	fmt.Println(MultiplicationTable(5))
}