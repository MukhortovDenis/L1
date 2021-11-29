package main

import (
	"fmt"
	"strings"
)

func main() {
	var str string = "snow dog sun"
	s := strings.Split(str, " ")
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	fmt.Println(strings.Join(s, " "))
}
