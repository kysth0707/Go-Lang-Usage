package main

import "fmt"

func main() {
	DictObj := make(map[string]string)
	DictObj["인삿말"] = "안녕하세요"
	DictObj["종료"] = "종료합니다람쥐"

	fmt.Println(DictObj["인삿말"])
	fmt.Println(DictObj["종료"])
}
