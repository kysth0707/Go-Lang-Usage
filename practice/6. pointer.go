package main

import "fmt"

func main() {
	Wow := 10
	PointerWow := &Wow
	fmt.Println("Pointer : ", PointerWow)
	fmt.Println("Value : ", *PointerWow)
}
