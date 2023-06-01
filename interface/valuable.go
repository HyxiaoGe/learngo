package main

import "fmt"

type stockPosition struct {
	ticker     string
	shareCount float32
	count      float32
}

func (s *stockPosition) getValue() float32 {
	return s.shareCount * s.count
}

type car struct {
	make  string
	model string
	price float32
}

func (c *car) getValue() float32 {
	return c.price
}

type valuable interface {
	getValue() float32
}

func showValue(asset valuable) {
	fmt.Printf("Value of the asset is %f\n", asset.getValue())
}

func main() {
	var v valuable
	v = &stockPosition{"GOOD", 577.20, 4}
	showValue(v)

	v = &car{"BMW", "M3", 66500}
	showValue(v)
}
