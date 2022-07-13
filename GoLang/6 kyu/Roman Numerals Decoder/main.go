package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(Decode("XXIV"))
}

func Decode(roman string) int {
	nums := []string{"IV", "IX", "XL", "XC", "CD", "CM", "I", "V", "X", "L", "C", "D", "M"}
	romans := map[string]int{"IV": 4, "IX": 9, "XL": 40, "XC": 90, "CD": 400, "CM": 900, "I": 1, "V": 5, "X": 10, "L": 50, "C": 100, "D": 500, "M": 1000}
	sum := 0
	for _, num := range nums {
	  sum += strings.Count(roman, num)*romans[num]
	  roman = strings.ReplaceAll(roman, num, "")
	}
	return sum
  }