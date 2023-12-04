package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	i := 1

	for {
		x := rand.Intn(100)
		// Move the cursor to the beginning of the line
		fmt.Print("\r")

		fmt.Printf("%v", x)
		if x == 50 {
			fmt.Printf("\tIterated %v times.\n", i)
			break
		}

		i++
		time.Sleep(50 * time.Millisecond)
	}
}

// create a for loop that uses break statement
// ● increment or decrement the variable being checked as a condition in the loop
// curriculum item # 082-hands-on-exercise-27
