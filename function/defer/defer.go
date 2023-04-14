package main

import "fmt"

func a() int {
	i := 0
	defer fmt.Println(i)
	i++
	return i
}

func f() {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}

func main() {
	//var num = a()
	//fmt.Println(num)
	f()
}
