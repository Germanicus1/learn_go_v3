package main

import (
	"fmt"
)

func main() {
	m := make(map[string][]string, 0)

	m[`fleming_ian`] = []string{`steaks`, `cigars`, `espionage`}
	m[`bond_james`] = []string{`shaken, not stirred`, `martinis`, `fast cars`}
	m[`moneypenny_jenny`] = []string{`intelligence`, `literature`, `computer science`}
	m[`no_dr`] = []string{`cats`, `ice cream`, `sunsets`}

	for k, v := range m {
		fmt.Printf("%v - ", k)
		for _, v2 := range v {
			fmt.Printf("%v, ", v2)
		}
		fmt.Printf("\n")
	}

}

/*
Create a map with a key of TYPE string which is a person’s “last_first” name, and a value of
TYPE []string which stores their favorite things. Store three records in your map. Print out all of
the values, along with their index position in the slice.
key					value

`bond_james`		`shaken, not stirred`, `martinis`, `fast cars`
`moneypenny_jenny`	`intelligence`, `literature`, `computer science`
`no_dr`				`cats`, `ice cream`, `sunsets`

*/
