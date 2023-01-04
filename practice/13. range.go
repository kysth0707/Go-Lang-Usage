package main

import (
	"fmt"
)

func main() {
	sum := 0

	nums := []int{4, 5, 6}
	for i, num := range nums {
		sum += num
		fmt.Println("i : ", i, ", num : ", num, ", sum : ", sum)
	}

	for i, c := range "go" {
		fmt.Println(i, c, string(c))
	}
}
