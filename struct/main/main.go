package main

import (
	"fmt"
	"github.com/HyxiaoGe/learngo/struct/struct_pack"
)

func main() {

	struct1 := new(struct_pack.ExpStruct)
	struct1.Mi1 = 10
	struct1.Mf1 = 9.9

	fmt.Printf("Mi1 = %d\n", struct1.Mi1)
	fmt.Printf("Mf1 = %f\n", struct1.Mf1)

}
