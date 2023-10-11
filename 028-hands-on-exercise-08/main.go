package main

import (
	"fmt"
	"math"
)

func main() {
	maxInt := 1<<8 - 1
	maxUint := 1<<7 - 1

	fmt.Printf("Max int8 = %d\n", maxInt)
	fmt.Printf("Max uint8 = %d\n", maxUint)
	fmt.Println(math.Pi)
}

/*
write a program that declares two variables
● one variable to store a VALUE of TYPE int8
	○ assign to it the largest number possible, then print it
● one variable to store a VALUE of TYPE uint8
	○ assign to it the largest number possible, then print it
*/
