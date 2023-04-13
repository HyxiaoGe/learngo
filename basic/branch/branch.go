package main

import (
	"fmt"
	"io/ioutil"
)

func main() {

	const filename = "abc.txt"
	//content, err := ioutil.ReadFile(filename)
	//if err != nil {
	//	fmt.Println(err)
	//} else {
	//	fmt.Printf("%s\n", content)
	//}

	if content, err := ioutil.ReadFile(filename); err == nil {
		fmt.Printf("%s\n", content)
	} else {
		fmt.Println(err)
	}

}
