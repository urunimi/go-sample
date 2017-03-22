package main

import "fmt"

const (
	_  = iota
	KB = 1 << (iota * 10) // 1 << (1 * 10)
	MB = 1 << (iota * 10) // 1 << (0 * 10)
	GB = 1 << (iota * 10) // 1 << (3 * 10)
	TB = 1 << (iota * 10) // 1 << (3 * 10)
)

func main() {
	fmt.Printf("%b\t", KB)
	fmt.Printf("%d\n", KB)
	fmt.Printf("%b\t", MB)
	fmt.Printf("%d\n", MB)
	fmt.Printf("%b\t", GB)
	fmt.Printf("%d\n", GB)
	fmt.Printf("%b\t", TB)
	fmt.Printf("%d\n", TB)
}
