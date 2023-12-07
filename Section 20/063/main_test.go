// put this code into it's own file
// Test files live in the same package as the code they test.
// They are named with the following convention:
// `filename_test.go` where filename is the name
// of the file that contains the code you want to test.

package main

import "testing"

func TestAdd(t *testing.T) {
	total := Add(5, 5)
	if total != 10 {
		t.Errorf("Sum was incorrect, got: %d, want: %d.", total, 10)
	}
}

// to run your test
// at the terminal, aka bash, aka shell
// go test
