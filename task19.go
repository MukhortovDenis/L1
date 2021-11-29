package main

import "fmt"

func main() {
	var str string = "главрыба"
	strToRune := []rune(str)
	for i, j := 0, len(strToRune)-1; i < j; i, j = i+1, j-1 {
		strToRune[i], strToRune[j] = strToRune[j], strToRune[i]
	}
	fmt.Print(string(strToRune))
}
