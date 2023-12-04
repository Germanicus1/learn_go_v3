/*
write a program that uses the following:
- var for zero value
- short declaration operator
- multiple initializations
- var when specificity is required
● blank identifier
print items as necessary to make the program interesting
*/

package main

import "fmt"

func main() {

	var zeroValue int
	short := 10
	var first, second float64 = 20, 30
	var pi float64 = 3.14
	var surface float64 = first * first * pi

	fmt.Printf("\n%v\t%v\t%v\t%v\t%v\t%v\n\n", zeroValue, short, first, second, pi, surface)

}
