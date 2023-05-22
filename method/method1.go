package main

import "fmt"

type TwoInts struct {
	a, b int
}

func main() {
	two1 := new(TwoInts)
	two1.a = 10
	two1.b = 20

	fmt.Printf("The sum is: %d\n", two1.AddThem())
	fmt.Printf("Add them to the param: %d\n", two1.AddToParam(20))
}

func (tn *TwoInts) AddThem() int {
	return tn.a + tn.b
}

func (tn *TwoInts) AddToParam(param int) int {
	return tn.a + tn.b + param
}
