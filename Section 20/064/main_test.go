package main

import (
	"testing"
)

func TestAdd(t *testing.T) {
	total := add(10, 20)
	if total != 30 {
		t.Errorf("\nAdd function retruned incorrect result. Expected: %d Got: %d", 30, total)
	}

}

func TestSubtract(t *testing.T) {
	result := subtract(20, 10)
	if result != 10 {
		t.Errorf("\nSubstract function returned incorrect result. Expected: %d Got: %d", 10, result)
	}
}

func TestDoMath(t *testing.T) {
	output := doMath(10, 20, add)

	if output != 30 {
		t.Errorf("\ndoMath returned incorrect result. Expected: %d Got: %d", 30, output)
	}
}
