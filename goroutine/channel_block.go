package main

import (
	"fmt"
	"time"
)

func main() {

	ch1 := make(chan int)
	go suck(ch1)
	go pump(ch1)
	time.Sleep(time.Second * 10)
	//fmt.Println(<-ch1)

}

func suck(ch chan int) {
	for {
		fmt.Println(<-ch)
	}
}

func pump(ch chan int) {
	for i := 0; ; i++ {
		ch <- i
	}
}
