package main

import (
	"fmt"
	"reflect"
)

func main() {

	var x float64 = 3.14
	v := reflect.ValueOf(x)
	//v.SetFloat(1.0) //panic: reflect: reflect.Value.SetFloat using unaddressable value

	fmt.Println("settability of v:", v.CanSet())
	v = reflect.ValueOf(&x)
	fmt.Println("type of v:", v.Type())
	fmt.Println("settability of v:", v.CanSet())

	v = v.Elem()
	fmt.Println("settability of v:", v.CanSet())
	fmt.Println("The Elem of v is:", v)
	v.SetFloat(3.1415)
	fmt.Println(v.Interface())
	fmt.Println(v)
}
