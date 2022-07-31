package main

import "fmt"

func main() {
	fmt.Println(Hero(7, 3))
}

func Hero(bullets, dragons int) bool {
	return bullets > dragons*2
}