package main

import "fmt"

func main() {
	HelloWorld()
	fmt.Println(Say("나", "안녕"))
	a := Say("나", "안녕")
	fmt.Println(a)

	fmt.Println(Divide(5, 2))
	c, d := Divide(5, 0)
	fmt.Println(c, d)
}

func HelloWorld() {
	fmt.Println("ㅎㅇ요")
}

func Say(From string, Message string) bool {
	fmt.Println(From, Message)

	return true
}

func Divide(a, b int) (result int, success bool) {
	if b == 0 {
		result = 0
		success = false
		return
	}
	result = a / b
	success = true
	return
}

// func 함수명(매개변수) 반환 타입

// 함수명이 대문자로 시작: Public 함수
// 함수명이 소문자로 시작: Private 함수
