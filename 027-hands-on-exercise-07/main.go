package main

import "fmt"

func main() {
	a, b, c := 747, 911, 90210

	fmt.Printf("%v\t%b\t%#x\n", a, a, a)
	fmt.Printf("%v\t%b\t%#x\n", b, b, b)
	fmt.Printf("%v\t%b\t%#x\n", c, c, c)

}

/*
write a program that uses print verbs to show the following numbers
	● 747
	● 911
	● 90210
as
	● decimal
	● binary
	● hexadecimal
*/
