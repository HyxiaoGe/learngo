package main

var o = "G"

func main() {
	n()
	m()
	n()
}

func n() { print(o) }

func m() {
	o := "O"
	print(o)
}
