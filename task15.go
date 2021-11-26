package main

import "fmt"

var justString string

func someFunc() {
	v := "somestring"
	a := len(v)
	justString = v[:a]
	fmt.Print(justString)
}

func main() {
	someFunc()
}
