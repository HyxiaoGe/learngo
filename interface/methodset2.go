package main

import "fmt"

type Appender interface {
	Append(int)
}

type List []int

func (l List) Len() int { return len(l) }

func (l *List) Append(val int) {
	*l = append(*l, val)
}

func CountInto(a Appender, start, end int) {
	for i := start; i <= end; i++ {
		a.Append(i)
	}
}

type Lender interface {
	Len() int
}

func LongEnough(l Lender) bool {
	return l.Len()*10 > 2
}

func main() {

	var list List
	if LongEnough(list) {
		fmt.Printf("- lst is long enough\n")
	}
	plst := new(List)

	CountInto(plst, 1, 10)
	if LongEnough(plst) {
		// VALID: a *List can be dereferenced for the receiver
		fmt.Printf("- plst is long enough\n")
	}
}
