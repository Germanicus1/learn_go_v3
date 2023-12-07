package main

import "fmt"

var (
	a, b, c *string
	d       *int
)

func init() {
	p := "Drop by drop the bucket gets filled."
	q := "Persitently, patiently, you're bound to suceed."
	r := "The meaning of life is..."
	n := 42
	a = &p
	b = &q
	c = &r
	d = &n
}

func main() {
	fmt.Printf("The value of variable a is %v and it is of type %T\n", a, a)
	fmt.Printf("The value of variable b is %v and it is of type %T\n", b, b)
	fmt.Printf("The value of variable c is %v and it is of type %T\n", c, c)
	fmt.Printf("The value of variable d is %v and it is of type %T\n", d, d)

	fmt.Printf("The value stored at adress %v and is %v\n", a, *a)
	fmt.Printf("The value stored at adress %v and is %v\n", b, *b)
	fmt.Printf("The value stored at adress %v and is %v\n", c, *c)
	fmt.Printf("The value stored at adress %v and is %v\n", d, *d)

}

/*
In the provided code, there are variables that store VALUES of TYPE *int and TYPE *string. The
values stored are memory addresses. Using the variables provided, do the following:
● print the VALUE stored in each variable
	○ these will be memory addresses
● print the TYPE of each variable
● print the data stored at memory locations
	○ dereference the stored memory address *
*/
