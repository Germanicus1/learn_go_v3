package main

import "fmt"

func main() {
	fmt.Println(printSquare(square, 4))
}

func square(n int) int {
	return n * n
}

func printSquare(f func(int) int, n int) string {

	x := f(n)
	return fmt.Sprintf("The square of %v is %v", n, x)
}

/*
A “callback” is when we pass a func into a func as an argument. For this exercise,
	● pass a func into a func as an argument
		○ func square(n int) int
		○ printSquare(f func(int)int, int) string

*/
