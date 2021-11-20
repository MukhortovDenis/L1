package main

import "fmt"

type Human struct {
	Name string
	Age  string
	Action
}
type Action struct {
	Status string
}

func (a *Action) showStatus() {
	fmt.Print("На данный момент " + a.Status + " !")
}
func main() {
	human := new(Human)
	human.Status = "жду"
	human.showStatus()
}
