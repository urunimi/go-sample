package main

import (
	"errors"
	"fmt"
	"log"
	"math/rand"
)

func main() {
	for i := 0; i < 100; i++ {
		doSomething()
	}
}

func Recover(randSize int) {
	if r := recover(); r != nil {
		if randSize > 0 && rand.Intn(randSize) == 0 {
			log.Fatal(errors.New(fmt.Sprintf("%s", r)))
		} else {
			log.Print(r)
		}
	}
}

func doSomething() {
	defer Recover(100)

	nextInt := rand.Intn(5)

	println("nextInt: ", nextInt)

	if nextInt == 0 {
		panic("panic!")
	}
}
