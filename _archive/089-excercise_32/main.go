package main

import (
	"fmt"
)

func main() {
	m := map[string]int{
		"James":      42,
		"Moneypenny": 32,
	}

	// age := m["James"]

	if age, ok := m["James"]; ok {
		fmt.Printf("Bond is %v years old\n", age)
	}

	// age = m["Q"]
	if _, ok := m["Q"]; !ok {
		fmt.Println("There is no Q")
	}

}

// ● add this code to print a single value stored in the map
// 		age := m["James"]
// 		fmt.Println(age)
// ● now use similar code to use the lookup of "Q" and print that value
// ● now use the "comma ok" idiom to test whether "Q" is actually stored in the map, then
//   print out a statement if it is not stored in the map
// 		○ hint: check effective go for the "comma ok" idiom
