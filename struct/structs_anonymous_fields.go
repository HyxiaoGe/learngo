package main

import "fmt"

type innerS struct {
	in1 int
	in2 int
}

type OuterS struct {
	b int
	c float32
	int
	innerS
}

func main() {

	outer := new(OuterS)
	outer.b = 6
	outer.c = 7.5
	outer.int = 60
	outer.in1 = 5
	outer.in2 = 10

	fmt.Printf("outer.b is: %d\n", outer.b)
	fmt.Printf("outer.c is: %f\n", outer.c)
	fmt.Printf("outer.int is: %d\n", outer.int)
	fmt.Printf("outer.in1 is: %d\n", outer.in1)
	fmt.Printf("outer.in2 is: %d\n", outer.in2)

	//	使用结构体字面量
	outer2 := OuterS{6, 7.5, 60, innerS{10, 20}}
	fmt.Println("outer2 is:", outer2)
}
