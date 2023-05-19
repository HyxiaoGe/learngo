package main

import "fmt"

func main() {

	s := make([]byte, 5)
	fmt.Println(len(s), cap(s))
	s = s[2:4]
	fmt.Println(len(s), cap(s))

	//s1 := []string{"p", "o", "e", "m"}
	//s2 := s1[2:]

	//s2[1] = "a"
	//fmt.Println(s1)
	//fmt.Println(s2)

}
