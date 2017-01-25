package main

// file level scope
import "fmt"

// package level scope
var x int = 42

func main() {
	fmt.Println(x)
	foo()
	// block level scope
	y := 17
	fmt.Println(y)
}

func foo() {
	fmt.Println(x)
}
