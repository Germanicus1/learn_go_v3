package main

import "fmt"

func main() {
	fmt.Println(Add(3, 4))
}

// Add.
// Adds the two arguments.
// @param	int, int	The two numbers we want to add together
// @return	int			The result of the addition in the form of an integer.
func Add(a int, b int) int {
	return a + b
}

// put this code into it's own file
// Test files live in the same package as the code they test.
// They are named with the following convention:
// `filename_test.go` where filename is the name
// of the file that contains the code you want to test.

/*
package main

import "testing"

func TestAdd(t *testing.T) {
    total := Add(5, 5)
    if total != 10 {
        t.Errorf("Sum was incorrect, got: %d, want: %d.", total, 10)
    }
}

*/

// to run your test
// at the terminal, aka bash, aka shell
// go test
