package main

import "fmt"

type IDuck interface {
	Quack()
	Walk()
}

type Bird struct{}

func (b Bird) Quack() {
	fmt.Println("I am quacking!")
}

func (b Bird) Walk() {
	fmt.Println("I am walking!")
}

func DuckDance(duck IDuck) {
	for i := 0; i < 3; i++ {
		duck.Quack()
		duck.Walk()
	}
}

func main() {
	b := new(Bird)
	DuckDance(b)
}
