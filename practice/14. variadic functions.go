package main

import "fmt"

func main() {
	fmt.Println(sumOf(10, 9, 8))

	nums := []int{1, 5, 9, 7}
	fmt.Println(sumOf(nums...))
}

func sumOf(nums ...int) int {
	fmt.Println(nums, ",")
	sum := 0

	for _, num := range nums {
		sum += num
	}
	return sum
}
