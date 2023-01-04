package main

import (
	"fmt"
	"math"
)

func main() {
	const a = 10000

	const b = 3e20 / 3e10

	fmt.Println(a)
	fmt.Println(int64(b))

	fmt.Println(math.Sin(a))
}
