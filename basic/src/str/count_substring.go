package main

import (
	"fmt"
	"strings"
)

func main() {
	var strTxt string = "Hello, how is it going, Hugo?"
	var manyG = "gggggggggg"

	fmt.Printf("%d\n", strings.Count(strTxt, "H"))
	fmt.Printf("%d\n", strings.Count(manyG, "g"))
	fmt.Printf("%s\n", strings.ToLower(strTxt))
	fmt.Printf("%s\n", strings.ToUpper(strTxt))
}
