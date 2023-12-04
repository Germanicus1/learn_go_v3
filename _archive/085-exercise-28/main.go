package main

import "fmt"

func main() {

	count := 100
	for i := 0; i < count; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Printf("%v\n", i)
	}

}

// use modulus and the continue statement in a loop to print all ODD numbers
// ● joke about the programmer and infinite loops
// curriculum item # 083-hands-on-exercise-28
