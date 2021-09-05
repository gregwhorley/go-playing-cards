package assignment1

import "fmt"

func main() {
	sliceOfInts := make([]int, 10)
	for i := range sliceOfInts {sliceOfInts[i] = i}
	for number := range sliceOfInts {
		if number % 2 == 0 {
			fmt.Printf("%d is even", number)
		} else {
			fmt.Printf("%d is odd", number)
		}
	}
}
