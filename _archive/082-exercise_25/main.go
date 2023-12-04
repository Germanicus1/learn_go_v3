package main

import (
	"fmt"
	"math/rand"
)

func main() {

	count := 42
	for i := 0; i < count; i++ {
		fmt.Printf("Iteration: %v\t", i+1)

		x := rand.Intn(5)

		switch x {
		case 0:
			fmt.Println("The variable contains the value 0.")
		case 1:
			fmt.Println("The variable contains the value 1.")
		case 2:
			fmt.Println("The variable contains the value 2.")
		case 3:
			fmt.Println("The variable contains the value 3.")
		case 4:
			fmt.Println("The variable contains the value 4.")
		}

	}
}

// ● Create one random int between 0 inclusive and 5 exclusive
// 		○ assign the value to a variable with the identifier x
// ● Use a switch statement to print a description of the variable and value
// ● run the code 42 times and print the iteration number
// curriculum item # 080-hands-on-exercise-25
