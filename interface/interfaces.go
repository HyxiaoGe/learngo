package main

import "fmt"

type Shaper interface {
	Area() float32
}

type Square struct {
	side float32
}

type Rectangle struct {
	length, width float32
}

func (square *Square) Area() float32 { return square.side * square.side }

func (rectangle *Rectangle) Area() float32 { return rectangle.length * rectangle.width }

func main() {

	square := &Square{5}
	rectangle := &Rectangle{5, 3}

	shapers := []Shaper{square, rectangle}
	fmt.Println("Looping through shapes for area ...")
	for _, shaper := range shapers {
		fmt.Println("Shape details: ", shaper)
		fmt.Println("Area of this shape is: ", shaper.Area())
	}
}
