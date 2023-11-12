package main

import (
	"fmt"
	"math/rand"
)

func main() {

	for i := 0; i < 100; i++ {
		if x := (rand.Intn(10)); x == 3 {
			fmt.Printf("count is %v\n", i)
			fmt.Println("x is 3")
			break
		}
	}

}

// use the "statement statement" idiom to
// 		○ initialize x with and random int between 0 inclusive and 5 exclusive
// 		○ if x is 3
// 			■ print "x is 3"
// ● run that code 100 times
// ● what's the benefit of using the "statement statement" idiom?
