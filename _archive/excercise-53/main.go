package main

import "fmt"

func main() {

	fmt.Println(foo(20))
	fmt.Println(bar())

}

func foo(x int) int {
	x += 1
	return x

}

func bar() (int, string) {
	return 42, "Hello world!"
}

/*
○ create a func with the identifier foo that returns an int
○ create a func with the identifier bar that returns an int and a string
○ call both funcs
○ print out their results

*/
