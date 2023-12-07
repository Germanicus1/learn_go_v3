package main

import "fmt"

func main() {

	x := returnFunc()
	fmt.Printf("%T\n", x)
	fmt.Println(x())
}

func returnFunc() func() int {
	return fortyTwo
}

func fortyTwo() int {
	return 42
}

/*
Create a func
	○ which returns a func
		■ which returns 42
		● assign the returned func to a variable
		● call the returned func
		● print

*/
