package main

import (
	"fmt"
	"reflect"
)

func main() {
	num := 10
	switch num {
	case 10:
		fmt.Println("10이다")
	default:
		fmt.Println("아니다")
	}

	fmt.Println("타입", reflect.TypeOf(num))
}
