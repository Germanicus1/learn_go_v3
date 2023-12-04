package main

import "fmt"

func main() {

	count := 5

	for i := 0; i < count; i++ {
		fmt.Printf("Outer loop run %v\n", i+1)
		for j := 0; j < count; j++ {
			fmt.Printf("\tInner loop run %v\n", j+1)
		}
	}

}

// create a loop that runs 5 times
// ● nest a loop within the first loop that also prints 5 times
// ● print something in each loop to illustrate what is occuring
