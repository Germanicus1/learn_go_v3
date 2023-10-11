package main

import "fmt"

type ByteSize int

const (
	_           = iota
	KB ByteSize = 1 << (10 * iota)
	MB
	GB
	TB
	PB
	EB
)

func main() {
	fmt.Printf("KB = %d Bytes\t\t\t\t%b\n", KB, KB)
	fmt.Printf("KB = %d Bytes\t\t\t%b\n", MB, MB)
	fmt.Printf("KB = %d Bytes\t\t\t%b\n", GB, GB)
	fmt.Printf("KB = %d Bytes\t\t%b\n", TB, TB)
	fmt.Printf("KB = %d Bytes\t\t%b\n", PB, PB)
	fmt.Printf("KB = %d Bytes\t\t%b\n", EB, EB)

}
