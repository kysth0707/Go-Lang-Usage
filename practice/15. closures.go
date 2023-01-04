package main

import (
	"fmt"
)

func main() {
	nextInt := intPlease()
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())
}

func intPlease() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}
