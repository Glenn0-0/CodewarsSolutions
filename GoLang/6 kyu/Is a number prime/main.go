package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(IsPrime(7))
}

func IsPrime(n int) bool {
  if n <= 1 {
    return false
  }
  if (n == 2 || n == 3) {
    return true
  }
  if (n % 2 == 0 || n % 3 == 0) {
    return false
  }
  
  for i := 5; i <= int(math.Sqrt(float64(n))); i += 6 {
    if (n % i == 0 || n % (i + 2) == 0) {
      return  false
    }
  }
  
  return true
}