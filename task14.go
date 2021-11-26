package main

import "fmt"

func main() {
	var a interface{} = make(chan int)
	switch v := a.(type) {
	case int:
		fmt.Print("Это инт ", v)

	case string:
		fmt.Print("Это строка ", v)

	case bool:
		fmt.Print("Это булево значение " , v)
	
	case chan int:
		fmt.Print("Это канал ", v)
	}
}
