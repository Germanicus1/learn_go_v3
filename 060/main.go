package main

import "fmt"

func main() {
	defer def1()

	fmt.Println("Deferred fuynctions:")

}

func def1() {
	defer def2()
	fmt.Println("First deferred function")
}
func def2() {
	defer def3()
	fmt.Println("Second deferred function")
}
func def3() {
	fmt.Println("Third deferred function")
}

/*
“defer” multiple functions in main
○ show that a deferred func runs after the func containing it exits.
○ determine the order in which the multiple defer funcs run

Last in First out
*/
