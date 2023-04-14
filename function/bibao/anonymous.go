package main

import "fmt"

func main() {

	//fplus := func(x, y int) int { return x + y }
	//fmt.Println(fplus(1, 2))

	fplus := func(x, y int) int { return x + y }(1, 2)
	fmt.Println(fplus)

}
