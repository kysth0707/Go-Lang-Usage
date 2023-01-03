package main

import (
	"fmt"
	"time"
)

func main() {
	go sayHello()
	sayHello2()
}

func sayHello() {
	for i := 0; i < 100; i++ {
		fmt.Println("안녕", i)
		time.Sleep(time.Second * 1)
	}
}

func sayHello2() {
	for i := 0; i < 100; i++ {
		fmt.Println("안녕22", i)
		time.Sleep(time.Second * 1)
	}
}
