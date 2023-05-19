package main

import "fmt"

var Days = map[int]string{1: "monday", 2: "tuesday", 3: "wednesday", 4: "thursday", 5: "friday", 6: "saturday", 7: "sunday"}

func main() {
	fmt.Println(Days)
	flagHolliday := false
	for k, v := range Days {
		if v == "thursday" || v == "holiday" {
			fmt.Println(v, " is the ", k, "th day in the week")
			if v == "holiday" {
				flagHolliday = true
			}
		}
	}
	if !flagHolliday {
		fmt.Println("holiday is not a day!")
	}

}
