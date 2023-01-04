package main

import (
	"fmt"
)

func main() {
	s := make([]string, 3)
	fmt.Println("emp : ", s)

	s[0] = "a"
	fmt.Println(s)
	fmt.Println(s[0])
	fmt.Println("len", len(s))

	s = append(s, "b")
	fmt.Println("append : ", s)

	fmt.Println("len : ", len(s))
}
