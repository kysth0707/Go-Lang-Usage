package main

import (
	"fmt"
)

type car struct {
	name  string
	speed int
}

func newCar(name string) *car {
	c := car{name: name}
	c.speed = 10
	return &c
}

func main() {
	fmt.Println(car{"a", 10})
	fmt.Println(car{name: "b", speed: 20})
	fmt.Println(&car{name: "c", speed: 20})
	fmt.Println(newCar("ㅎㅇㅎㅇ"))

	myCar := newCar("ㅎㅇㅎㅇ22")
	fmt.Println(myCar)
	myCar.name = "이름바꿈"
	fmt.Println(myCar)

}
