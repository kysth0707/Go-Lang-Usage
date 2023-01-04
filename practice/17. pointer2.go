package main

import (
	"fmt"
)

func toZeroVal(val int) {
	val = 0
}

func toZeroPtr(addr *int) {
	*addr = 0
}

func main() {
	i := 1
	fmt.Println("i : ", i)
	toZeroVal(i)
	fmt.Println("i : ", i)

	i = 1
	fmt.Println("i : ", i)
	toZeroPtr(&i)
	fmt.Println("i : ", i)
}
