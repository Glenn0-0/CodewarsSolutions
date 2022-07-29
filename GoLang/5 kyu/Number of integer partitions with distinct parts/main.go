package main

import (
	"fmt"
	"math/big"
)

func main() {
	fmt.Println(Partitions(5))
}

func Partitions(n int) *big.Int {
	arr := []int{}
	for i := 0; i <= n; i++ {
	  if i == 0 {
		arr = append(arr, 1)
	  } else {
		arr = append(arr, 0)
	  }
	}
  
	for k := 1; k < n; k++ {
	  temp := []int{}
	  for i := 0; i <= n; i++ {
		if (i - k) < 0 {
		  temp = append(temp, arr[i])
		} else {
		  temp = append(temp, arr[i] + arr[i-k])
		}
	  }
	  arr = temp
	}
	return big.NewInt(int64(arr[n] + 1))
}