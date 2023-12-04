package main

import (
	"fmt"
)

func main() {
	// x := make([][]string, 0, 2)

	x := [][]string{{"James", "Bond", "Shaken, not stirred"}, {"Miss", "Moneypenny", "I'm 008."}}

	for i, x := range x {
		fmt.Println(i, x)
		for a, b := range x {
			fmt.Println(a, b)
		}
	}

}

/*
Create a slice of a slice of string ([][]string). Store the following data in the multi-dimensional
slice:
"James", "Bond", "Shaken, not stirred"
"Miss", "Moneypenny", "I'm 008."
Range over the records, then range over the data in each record.
*/
