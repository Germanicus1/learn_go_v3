package main

import "fmt"

var name string = "Peter"

const height int = 172

func main() {
	surname := "Kerschbaumer"
	const weight int = 76

	fmt.Printf("My name is %s %s. I'm %d cm tall and I weigh %d kg.\n", name, surname, height, weight)

}
