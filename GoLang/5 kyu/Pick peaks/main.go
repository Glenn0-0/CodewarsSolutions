package main

import "fmt"

type PosPeaks struct {
	Pos   []int
	Peaks []int
}

func main() {
	fmt.Println(PickPeaks([]int{3, 2, 3, 6, 4, 1, 2, 3, 2, 1, 2, 3}))
}
  
func goesLower(arr []int, ind int) bool {
	for k := ind; k < len(arr); k++ {
	  
	  if arr[k] > arr[ind] {
		return false
		
	  } else if arr[k] < arr[ind] {
		return true
	  }
	  
	}
	
	return false
}
  
func PickPeaks(array []int) PosPeaks {
	var pos, peaks = []int{}, []int{}
	for i := 1; i < len(array) - 1; i++ {
	  
	  if array[i] > array[i - 1] && array[i] >= array[i + 1] && goesLower(array, i) {
		pos = append(pos, i)
		peaks = append(peaks, array[i])
	  }
	  
	}
	
	result := PosPeaks {
	  Pos: pos,
	  Peaks: peaks,
	}
	
	return result
}