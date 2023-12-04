package main

import "fmt"

func main() {
	x := [5]int{1, 2, 3, 4, 5}

	for i := range x {
		fmt.Printf("%v - %T\n", x[i], x[i])
	}

}

// Using a COMPOSITE LITERAL:
// ● create an ARRAY which holds 5 VALUES of TYPE int
// ● assign VALUES to each index position.
// ● Range over the array and print the values out.
// ○ print out the VALUE a
