package main

import (
	"math"
	"fmt"
)

func main() {
	fmt.Println(F(1000))
}

func isPrime(n int) bool {
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

func countEvens(n int) (count int) {
  for n > 0 {
    if (n%10) % 2 == 0 {
      count++
    }
    n /= 10
  }
  return count
}

func countNums(n int) (count int) {
  for n > 9 {
    count++
    n /= 10
  }
  count++
  if n == 1 {
    count--
  }
  return count
}

func F(n int) int {
  for i := n - 1; i > 0; i-- {
    if isPrime(i) && countEvens(i) == (countNums(i) - 1) {
    return i
    }
  }
  return 0
}