package main

import "fmt"

func main() {
	fmt.Println(TwiceAsOld(29, 7))
}

func TwiceAsOld(dadYearsOld, sonYearsOld int) int {
	for i := 0; i < 200; i++ {
	  if (dadYearsOld - i) == (sonYearsOld - i)*2 || (dadYearsOld + i) == (sonYearsOld + i)*2 {
		return i
	  }
	}
	return 0
  }