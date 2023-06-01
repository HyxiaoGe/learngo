package main

import (
	"fmt"
	"math"
)

/**
类型断言
*/
type Shaper1 interface {
	Area() float32
}

type Square1 struct {
	side float32
}

func (s *Square1) Area() float32 { return s.side * s.side }

type Circle struct {
	radius float32
}

func (c *Circle) Area() float32 { return c.radius * c.radius * math.Pi }

func main() {

	var areaIntf Shaper1
	square1 := new(Square1)
	square1.side = 5

	areaIntf = square1
	if s, ok := areaIntf.(*Square1); ok {
		fmt.Printf("The type of areaIntf is: %T\n", s)
	}

	circle1 := new(Circle)
	circle1.radius = 2.5
	areaIntf = circle1
	if c, ok := areaIntf.(*Circle); ok {
		fmt.Printf("The type of areaIntf is: %T\n", c)
	} else {
		fmt.Println("areaIntf does not contain a variable of type Circle")
	}

	switch t := areaIntf.(type) {
	case *Square:
		fmt.Printf("Type Square %T with value %v\n", t, t)
	case *Circle:
		fmt.Printf("Type Circle %T with value %v\n", t, t)
	case nil:
		fmt.Printf("nil value: nothing to check?\n")
	default:
		fmt.Printf("Unexpected type %T\n", t)
	}

}
