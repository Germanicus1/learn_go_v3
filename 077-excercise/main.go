package main

import (
	"fmt"
	"math/rand"
)

func init() {
	fmt.Println("My program has initialized correctly.")
}

func main() {

	x := (rand.Intn(250))

	fmt.Printf("The value of x is %v and its type is %T\n", x, x)

	switch {
	case x >= 0 && x <= 100:
		fmt.Println("x is between 0 and 100")
	case x >= 101 && x <= 200:
		fmt.Println("x is between 101 and 200")
	case x >= 201 && x <= 250:
		fmt.Println("x is between 201 and 250")
	}

	count := 100
	for i := 0; i < count; i++ {
		x := rand.Intn(3)
		if x == 3 {
			fmt.Println(x)
		} else if x != 3 && i == count-1 {
			fmt.Println("No occurance of the number 3.")
		}
	}

}

/*
create a program that uses both SEQUENTIAL and CONDITIONAL control flow. Your
program should do the following
○ create a random int between 0 and 250
○ store the value of that int in a variable with the identifier of x
■ func Intn(n int) int
● rand.Intn()
○ print out the name and value of the variable
Todd McLeod - Learn To Code Go - Page 50
○ use an if statement to do the following
■ if the value is between 0 and 100
● print between 0 and 100
■ if the value is between 101 and 200
● print between 101 and 200
■ if the value is between 201 and 250
● print between 201 and 250
● re: inclusive, exclusive – does rand.Intn()
○ include zero in its output?
○ include 250 in its output?
○ show this in code using the numbers 0 - 3
*/
