package main

import "fmt"

func wrapper(x int) func() int {
	return func() int {
		x++
		return x
	}
}

func main() {
	increment := wrapper(0)

	fmt.Println(increment())
	fmt.Println(increment())

	increment = wrapper(1)

	fmt.Println(increment())
	fmt.Println(increment())
}
